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

func PrintParams() {
	for _, s := range os.Args[1:] {
		printStrEndl(s)
	}
}
