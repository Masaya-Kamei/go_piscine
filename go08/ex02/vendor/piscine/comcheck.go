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
	checkMap := map[string]struct{}{"42": {}, "piscine": {}, "piscine 42": {}}

	for _, arg := range os.Args[1:] {
		if _, ok := checkMap[arg]; ok {
			printStrEndl("Alert!!!")
			return
		}
	}
}
