package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	start := 0
	for i, r := range os.Args[0] {
		if r == '/' {
			start = i + 1
		}
	}
	for _, r := range []rune(os.Args[0])[start:] {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
