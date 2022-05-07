package main

import (
	"ft"
	"os"
)

type boolean int

const (
	no      boolean = 0
	yes     boolean = 1
	EvenMsg         = "I have an even number of arguments"
	OddMsg          = "I have an odd number of arguments"
)

var lengthOfArg int = getArgc(os.Args)

func getArgc(args []string) int {
	argc := 0
	for range args {
		argc++
	}
	return argc
}

func even(nbr int) int { return nbr % 2 }

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
