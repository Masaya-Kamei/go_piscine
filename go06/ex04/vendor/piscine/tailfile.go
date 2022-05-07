package piscine

import (
	"os"
)

func getBytesLen(bytes []byte) int {
	i := 0
	for range bytes {
		i++
	}
	return i
}

func tailLines(bytes []byte) {
	nlCount := 0
	startIndex := getBytesLen(bytes) - 1
	if startIndex >= 0 && bytes[startIndex] != '\n' {
		nlCount++
	}
	for startIndex >= 0 {
		if bytes[startIndex] == '\n' {
			nlCount++
			if nlCount == 11 {
				break
			}
		}
		startIndex--
	}
	printStr(string(bytes[startIndex+1:]))
}

func tailBytes(bytes []byte, optionC int) {
	startIndex := getBytesLen(bytes) - optionC
	if startIndex < 0 {
		startIndex = 0
	}
	printStr(string(bytes[startIndex:]))
}

func tailFile(fileName string, optionC int, printFileName func(s string)) int {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		printStrEndl(err.Error())
		return 1
	}

	if printFileName != nil {
		printFileName(fileName)
	}

	if optionC == -1 {
		tailLines(bytes)
	} else {
		tailBytes(bytes, optionC)
	}
	return 0
}
