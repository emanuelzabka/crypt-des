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

func readArgs() {
	var encrypt, decrypt, newKey, help bool
	flag.BoolVar(&encrypt, "encrypt", false, "Encrypt input")
	flag.BoolVar(&decrypt, "decrypt", false, "Decrypt input")
	flag.BoolVar(&newKey, "newkey", false, "Outputs a new key")
	flag.BoolVar(&help, "h", false, "Display help and exit")
	flag.StringVar(&inputParam, "i", "", "Input file to encrypt/decrypt. Use \"-\" to standard input. Default \"-\"")
	flag.StringVar(&outputParam, "o", "", "Output file. Use \"-\" to standard output. Default \"-\"")
	flag.StringVar(&keyString, "k", "", "Cipher key")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if newKey {
		key := prepareKey()
		fmt.Println(blockToHexString(key))
		os.Exit(0)
	}
	if !encrypt && !decrypt || encrypt {
		operation = des.ENCRYPT
	} else {
		operation = des.DECRYPT
	}
	if keyString == "" && operation == des.DECRYPT {
		fmt.Fprintln(os.Stderr, "** Cipher key is required for operation decrypt")
		flag.Usage()
		os.Exit(1)
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
}

func validateArgs() {
	// Input is readable
	if inputParam != "-" {
		if _, err := os.Stat(inputParam); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "Error: Input file not found")
		}
	}
}

func prepareKey() (result []byte) {
	if keyString == "" {
		rand.Seed(time.Now().UnixNano())
		result = make([]byte, BLOCK_SIZE)
		for i := range result {
			result[i] = byte(rand.Intn(256))
		}
	} else {
		keyString = fmt.Sprintf("%016s", keyString)
		keyString = keyString[0:16]
		var err error
		result, err = hex.DecodeString(keyString)
		if err != nil {
			fmt.Printf("Key %s is not a valid hexadecimal value", keyString)
			os.Exit(1)
		}
	}
	return result
}

func blockToHexString(block []byte) string {
	return hex.EncodeToString(block)
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

func encrypt(key []byte) {
	roundKeys := des.GenerateRoundsKeys(key, des.ENCRYPT)
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

func decrypt(key []byte) {
	end := false
	size := 0
	pad := 0
	roundKeys := des.GenerateRoundsKeys(key, des.DECRYPT)
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
	key := prepareKey()
	if keyString == "" {
		fmt.Fprintf(os.Stderr, "Using key: %s\n", blockToHexString(key))
	}
	if operation == des.ENCRYPT {
		encrypt(key)
	} else {
		decrypt(key)
	}
}
