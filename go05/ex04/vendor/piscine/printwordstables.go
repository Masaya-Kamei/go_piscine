package piscine

import "ft"

func printStrEndl(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	for _, w := range a {
		printStrEndl(w)
	}
}
