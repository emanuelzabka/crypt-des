package main

import (
	"./des"
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
	"strings"
)

const BLOCK_SIZE = 8

var inputReader *bufio.Reader
var inputFile *os.File
var outputWriter *bufio.Writer
var outputFile *os.File
var inputParam string
var outputParam string
var keyString string
var operation int
var tripleDes bool

func askForKey() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, "Enter key: ")
	key, _ := reader.ReadString('\n')
	return strings.Trim(key, "\n")
}

func readArgs() {
	var encrypt, decrypt, newKey, help bool
	flag.BoolVar(&encrypt, "encrypt", false, "Encrypt input")
	flag.BoolVar(&decrypt, "decrypt", false, "Decrypt input")
	flag.BoolVar(&newKey, "newkey", false, "Outputs a new key")
	flag.BoolVar(&help, "h", false, "Display help and exit")
	flag.BoolVar(&tripleDes, "3des", false, "Encrypt/decrypt using Triple DES")
	flag.StringVar(&inputParam, "i", "", "Input file to encrypt/decrypt. Use \"-\" to standard input. Default \"-\"")
	flag.StringVar(&outputParam, "o", "", "Output file. Use \"-\" to standard output. Default \"-\"")
	flag.StringVar(&keyString, "k", "", "Cipher key")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if newKey {
		keys := prepareKey()
		fmt.Println(keysToString(keys))
		os.Exit(0)
	}
	if !encrypt && !decrypt || encrypt {
		operation = des.ENCRYPT
	} else {
		operation = des.DECRYPT
	}
	if encrypt && decrypt {
		fmt.Fprintln(os.Stderr, "Choose between only one of -encrypt or -decrypt operations")
		os.Exit(1)
	}
	if inputParam == "" {
		inputParam = "-"
	}
	if outputParam == "" {
		outputParam = "-"
	}
	if keyString == "" && operation == des.DECRYPT {
		// Request key from user if input is not standard input
		if inputParam != "-" {
			keyString = askForKey()
		}
		if keyString == "" {
			fmt.Fprintln(os.Stderr, "** Cipher key is required for operation decrypt")
			flag.Usage()
			os.Exit(1)
		}
	}
}

func validateArgs() {
	// Input is readable
	if inputParam != "-" {
		if _, err := os.Stat(inputParam); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "Error: Input file not found")
		}
	}
}

func prepareKey() (result [][]byte) {
	keyCount := 1
	if tripleDes {
		keyCount = 3
	}
	if keyString == "" {
		result = make([][]byte, keyCount)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < keyCount; i++ {
			result[i] = make([]byte, BLOCK_SIZE)
			for j := range result[i] {
				result[i][j] = byte(rand.Intn(256))
			}
		}
	} else {
		splitedKey := strings.Split(keyString, ":")
		if len(splitedKey) < keyCount {
			fmt.Fprintln(os.Stderr, "Triple DES requires three keys")
			os.Exit(1)
		}
		if len(splitedKey) == 3 {
			tripleDes = true
			keyCount = 3
		}
		if len(splitedKey) != keyCount {
			fmt.Fprintln(os.Stderr, "Invalid key count. Allowed one or three keys")
			os.Exit(1)
		}
		result = make([][]byte, keyCount)
		for k := range splitedKey {
			key := fmt.Sprintf("%016s", splitedKey[k])
			key = key[0:16]
			var err error
			result[k], err = hex.DecodeString(key)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Key %s is not a valid hexadecimal value\n", key)
				os.Exit(1)
			}
		}
	}
	return result
}

func blockToHexString(block []byte) string {
	return hex.EncodeToString(block)
}

func keysToString(keys [][]byte) string {
	res := make([]string, len(keys))
	for k := range keys {
		res[k] = blockToHexString(keys[k])
	}
	return strings.Join(res, ":")
}

func checkInputReader() {
	if inputReader != nil {
		return
	}
	if inputParam != "-" {
		file, err := os.Open(inputParam)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %s\n", err.Error())
			os.Exit(1)
		}
		inputFile = file
		inputReader = bufio.NewReader(file)
	} else {
		inputReader = bufio.NewReader(os.Stdin)
	}
}

func getNextInputBlock(block []byte) (end bool, size int, pad int) {
	end = false
	pad = 0
	checkInputReader()
	size, err := inputReader.Read(block)
	if err == io.EOF {
		end = true
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %s\n", err.Error())
		os.Exit(1)
	}
	if size > 0 && size != BLOCK_SIZE {
		if operation == des.DECRYPT {
			pad = int(block[0])
			size = 0
		} else {
			pad = BLOCK_SIZE - size
		}
	}
	return end, size, pad
}

func checkOutputWriter() {
	if outputWriter != nil {
		return
	}
	if outputParam != "-" {
		file, err := os.OpenFile(outputParam, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file: %s\n", err.Error())
			os.Exit(1)
		}
		outputFile = file
		outputWriter = bufio.NewWriter(file)
	} else {
		outputWriter = bufio.NewWriter(os.Stdout)
	}
}

func writeToOutput(block []byte, pad int) {
	checkOutputWriter()
	var err error
	if pad == 0 {
		_, err = outputWriter.Write(block)
	} else {
		if operation == des.ENCRYPT {
			buff := make([]byte, BLOCK_SIZE+1)
			copy(buff, block)
			buff[len(buff)-1] = byte(pad)
			_, err = outputWriter.Write(buff)
		} else {
			buff := make([]byte, BLOCK_SIZE-pad)
			copy(buff, block)
			_, err = outputWriter.Write(buff)
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %s\n", err.Error())
		os.Exit(1)
	}
	if outputParam == "-" {
		outputWriter.Flush()
	}
}

func closeFiles() {
	if outputWriter != nil {
		outputWriter.Flush()
	}
	if inputFile != nil {
		inputFile.Close()
	}
	if outputFile != nil {
		outputFile.Close()
	}
}

func encrypt(keys [][]byte) {
	roundKeys := des.GenerateRoundsKeys(keys, des.ENCRYPT)
	end := false
	size := 0
	pad := 0
	block := make([]byte, BLOCK_SIZE)
	for !end {
		end, size, pad = getNextInputBlock(block)
		if size > 0 {
			processedBlock := des.CipherBlock(block, roundKeys)
			writeToOutput(processedBlock, pad)
		}
	}
}

func decrypt(keys [][]byte) {
	end := false
	size := 0
	pad := 0
	roundKeys := des.GenerateRoundsKeys(keys, des.DECRYPT)
	var pendingBlock bool = false
	block := make([]byte, BLOCK_SIZE)
	var processedBlock []byte
	for !end {
		end, size, pad = getNextInputBlock(block)
		if pendingBlock {
			writeToOutput(processedBlock, pad)
			pendingBlock = false
		}
		if size > 0 {
			processedBlock = des.CipherBlock(block, roundKeys)
			pendingBlock = true
		}
	}
	if pendingBlock {
		writeToOutput(processedBlock, pad)
	}
}

func main() {
	defer closeFiles()
	readArgs()
	validateArgs()
	keys := prepareKey()
	if keyString == "" {
		fmt.Fprintf(os.Stderr, "Using key: %s\n", keysToString(keys))
	}
	if operation == des.ENCRYPT {
		encrypt(keys)
	} else {
		decrypt(keys)
	}
}
