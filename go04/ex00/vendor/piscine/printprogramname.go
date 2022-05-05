package piscine

import (
	"ft"
	"os"
)

func printStrEndl(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func PrintProgramName() {
	start := 0
	for i, r := range os.Args[0] {
		if r == '/' {
			start = i + 1
		}
	}
	printStrEndl(os.Args[0][start:])
}
