package main

import (
	"ft"
	"piscine"
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	ft.PrintRune('\n')

	// piscine.PrintNbrBase(0, "ab")
	// ft.PrintRune('\n')
	// piscine.PrintNbrBase(-1, "ab")
	// ft.PrintRune('\n')
	// piscine.PrintNbrBase(125, "")
	// ft.PrintRune('\n')
	// piscine.PrintNbrBase(125, "a")
	// ft.PrintRune('\n')
	// piscine.PrintNbrBase(125, "01234567890")
	// ft.PrintRune('\n')
	// piscine.PrintNbrBase(125, "0123456789+")
	// ft.PrintRune('\n')
}
