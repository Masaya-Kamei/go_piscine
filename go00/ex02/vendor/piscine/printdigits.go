package piscine

import "ft"

func PrintDigits() {
	for n := '0'; n <= '9'; n++ {
		ft.PrintRune(n)
	}
	ft.PrintRune('\n')
}
