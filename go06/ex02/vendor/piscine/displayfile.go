package piscine

import (
	"os"
)

func DisplayFile() {
	if err := validateArgc(); err != nil {
		printStrEndl(err.Error())
		return
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		printStrEndl(err.Error())
	}
	printStr(string(bytes))
}
