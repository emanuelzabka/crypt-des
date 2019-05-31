package des

var initialPermutation []int = []int{
	58, 50, 42, 34, 26, 18, 10, 2, 60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6, 64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1, 59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5, 63, 55, 47, 39, 31, 23, 15, 7,
}

var finalPermutation []int = []int{
	40, 8, 48, 16, 56, 24, 64, 32, 39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30, 37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28, 35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26, 33, 1, 41, 9, 49, 17, 57, 25,
}

var keyPermutation []int = []int{
	57, 49, 41, 33, 25, 17, 9, 1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27, 19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15, 7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29, 21, 13, 5, 28, 20, 12, 4,
}

var keyShiftsPerRound []int = []int{
	1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1,
}

var compressionPermutation []int = []int{
	14, 17, 11, 24, 1, 5, 3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8, 16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55, 30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53, 46, 42, 50, 36, 29, 32,
}

var expansionPermutation []int = []int{
	32, 1, 2, 3, 4, 5, 4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13, 12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21, 20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32, 1,
}

var sboxes [][]int = [][]int{
	{
		14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7,
		0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8,
		4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0,
		15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13,
	},
	{
		15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10,
		3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5,
		0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15,
		13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9,
	},
	{
		10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8,
		13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1,
		13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7,
		1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12,
	},
	{
		7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15,
		13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9,
		10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4,
		3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14,
	},
	{
		2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9,
		14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6,
		4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14,
		11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3,
	},
	{
		12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11,
		10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8,
		9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6,
		4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13,
	},
	{
		4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1,
		13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6,
		1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2,
		6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12,
	},
	{
		13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7,
		1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2,
		7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8,
		2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11,
	},
}

var pBoxPermutation []int = []int{
	16, 7, 20, 21, 29, 12, 28, 17, 1, 15, 23, 26, 5, 18, 31, 10,
	2, 8, 24, 14, 32, 27, 3, 9, 19, 13, 30, 6, 22, 11, 4, 25,
}

func permutate(block []byte, permutationMatrix []int) (result []byte) {
	result = make([]byte, len(permutationMatrix)>>3)
	for i := 0; i < len(permutationMatrix); i++ {
		resByteIndex := i >> 3
		resBitIndex := i - (resByteIndex << 3)
		blockByteIndex := (permutationMatrix[i] - 1) >> 3
		blockBitIndex := (permutationMatrix[i] - 1) - (blockByteIndex << 3)
		blockByte := block[blockByteIndex]
		resByte := result[resByteIndex]
		var bit byte = (0x01 & (blockByte >> uint(7-blockBitIndex))) << uint(7-resBitIndex)
		result[resByteIndex] = (resByte & ^(0x01 << uint(7-resBitIndex))) | bit
	}
	return result
}

func key64To56(key []byte) (result []byte) {
	result = permutate(key, keyPermutation)
	return result
}

func leftRoundShift(block []byte, shift int, beginOffset int, endOffset int) (result []byte) {
	result = make([]byte, len(block))
	copy(result, block)
	if endOffset > 0 {
		result[len(block)-1] = result[len(block)-1] & (0xff << uint(endOffset))
	}
	for i := 0; i < shift; i++ {
		var leading byte = 0x00
		for j := len(result) - 1; j >= 0; j-- {
			var nextLeading byte = (result[j] & 0x80) >> 7
			result[j] = (result[j] << 1) | leading
			leading = nextLeading
		}
		if endOffset > 0 {
			result[len(result)-1] = result[len(result)-1] | (leading << uint(endOffset))
		} else {
			result[len(result)-1] = result[len(result)-1] | leading
		}
	}
	return result
}

func joinKeys(left []byte, right []byte) (result []byte) {
	result = make([]byte, 7)
	for i := 0; i < 3; i++ {
		result[i] = left[i]
	}
	result[3] = (left[3] & 0xf0) | (0x0f & right[0])
	for i := 1; i < 4; i++ {
		result[3+i] = right[i]
	}
	return result
}

func getRoundSubkey(key []byte, round int) (result []byte) {
	var left []byte = make([]byte, 4)
	var right []byte = make([]byte, 4)
	copy(left, key[0:4])
	copy(right, key[3:])
	left = leftRoundShift(left, keyShiftsPerRound[round], 0, 4)
	right = leftRoundShift(right, keyShiftsPerRound[round], 4, 0)
	shiftedKey := joinKeys(left, right)
	result = permutate(shiftedKey, compressionPermutation)
	return result
}

func xorBlocks(block1 []byte, block2 []byte) (result []byte) {
	result = make([]byte, len(block1))
	for i := range block1 {
		result[i] = block1[i] ^ block2[i]
	}
	return result
}

func getBit(block []byte, pos int) (result byte) {
	posByte := pos >> 3
	posBit := pos - posByte<<3
	result = 0x01 & (block[posByte] >> (7 - uint(posBit)))
	return result
}

func getSBoxValue(block []byte, step int) (result byte) {
	beginPos := step * 6
	endPos := beginPos + 5
	var col byte
	var row byte

	row = (getBit(block, beginPos) << 1) | (getBit(block, endPos))
	col =
		(getBit(block, beginPos+1) << 3) |
			(getBit(block, beginPos+2) << 2) |
			(getBit(block, beginPos+3) << 1) |
			getBit(block, beginPos+4)
	result = byte(sboxes[step][row*16+col])
	return result
}

func applySBoxes(block []byte) (result []byte) {
	result = make([]byte, 4)
	for i := 0; i < 8; i += 2 {
		leftNibble := getSBoxValue(block, i)
		rightNibble := getSBoxValue(block, i+1)
		result[i>>1] = leftNibble<<4 | rightNibble
	}
	return result
}

func feistelFunction(key []byte, block []byte) (result []byte) {
	expandedBlock := permutate(block, expansionPermutation)
	expandedBlock = xorBlocks(expandedBlock, key)
	transformedBlock := applySBoxes(expandedBlock)
	result = permutate(transformedBlock, pBoxPermutation)
	return result
}

func GenerateRoundKeys(key []byte) (result [][]byte) {
	transformedKey := key64To56(key)
	result = make([][]byte, 16)
	for i := 0; i < 16; i++ {
		result[i] = getRoundSubkey(transformedKey, i)
	}
	return result
}

func EncryptBlock(block []byte, roundKeys [][]byte) (result []byte) {
	block = permutate(block, initialPermutation)
	left := block[0:4]
	right := block[4:]
	for k := range roundKeys {
		lastRight := right
		transformedRight := feistelFunction(roundKeys[k], right)
		right = xorBlocks(left, transformedRight)
		left = lastRight
	}
	result = permutate(append(right, left...), finalPermutation)
	return result
}

func DecryptBlock(block []byte, roundKeys [][]byte) (result []byte) {
	invertedKeys := make([][]byte, len(roundKeys))
	for i, j := len(roundKeys)-1, 0; i >= 0; i, j = i-1, j+1 {
		invertedKeys[j] = roundKeys[i]
	}
	return EncryptBlock(block, invertedKeys)
}
