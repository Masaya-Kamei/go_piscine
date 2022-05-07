package piscine

import "ft"

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func printStrEndl(s string) {
	printStr(s)
	ft.PrintRune('\n')
}
