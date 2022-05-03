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

func strsLen(strs []string) int {
	i := 0
	for range strs {
		i++
	}
	return i
}

func RevParams() {
	strs := os.Args[1:]
	len := strsLen(strs)
	for i := len - 1; i >= 0; i-- {
		printStrEndl(strs[i])
	}
}
