package main

import (
	"ft"
	"piscine"
)

func printStrEndl(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func main() {
	result := piscine.Rot14("Hello! How are You?")
	printStrEndl(result)

	ft.PrintRune('\n')
	printStrEndl(piscine.Rot14("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	printStrEndl(piscine.Rot14("abcdefghijklmnopqrstuvwxyz"))
	printStrEndl(piscine.Rot14("あいうえお"))
	printStrEndl(piscine.Rot14(""))
}
