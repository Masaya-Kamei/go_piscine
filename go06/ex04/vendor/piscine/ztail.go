package piscine

import (
	"ft"
	"os"
)

func getArgInfo(optionC int) ([]string, int) {
	args := os.Args[1:]
	if optionC != -1 {
		args = os.Args[3:]
	}
	argc := 0
	for range args {
		argc++
	}
	return args, argc
}

func fileNamePrinter(argc, successFlag int) func(fileName string) {
	return func(fileName string) {
		if argc >= 2 {
			if successFlag == 1 {
				ft.PrintRune('\n')
			}
			printStrEndl("==> " + fileName + " <==")
		}
	}
}

func ZTail() {
	optionC, err := readOptionC()
	if err != nil {
		printStrEndl(err.Error())
		os.Exit(1)
	}

	args, argc := getArgInfo(optionC)
	if argc == 0 {
		os.Exit(tailFile("/dev/stdin", optionC, nil))
	}
	exitStatus := 0
	successFlag := 0
	for _, fileName := range args {
		status := tailFile(fileName, optionC, fileNamePrinter(argc, successFlag))
		exitStatus |= status
		successFlag |= ^status & 1
	}
	os.Exit(exitStatus)
}
