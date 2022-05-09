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

func ComCheck() {
	checkStrs := []string{"42", "piscine", "piscine 42"}
	for _, arg := range os.Args[1:] {
		for _, s := range checkStrs {
			if arg == s {
				printStrEndl("Alert!!!")
				return
			}
		}
	}
}
