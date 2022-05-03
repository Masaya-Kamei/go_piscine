package main

import (
	"ft"
	"piscine"
)

func main() {
	ft.PrintRune(piscine.NRune("Hello!", 3))
	ft.PrintRune(piscine.NRune("Salut!", 2))
	ft.PrintRune(piscine.NRune("Bye!", -1))
	ft.PrintRune(piscine.NRune("Bye!", 5))
	ft.PrintRune(piscine.NRune("Ola!", 4))

	// ft.PrintRune(piscine.NRune("Hello!", 0))
	// ft.PrintRune(piscine.NRune("Hello!", 1))
	// ft.PrintRune(piscine.NRune("Hello!", 6))
	// ft.PrintRune(piscine.NRune("Hello!", 7))
	// ft.PrintRune(piscine.NRune("", 1))
	ft.PrintRune('\n')
}
