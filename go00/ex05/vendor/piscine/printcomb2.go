package piscine

import "ft"

func printNumber(nb int) {
	ft.PrintRune(rune(nb/10 + '0'))
	ft.PrintRune(rune(nb%10 + '0'))
}

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			printNumber(a)
			ft.PrintRune(' ')
			printNumber(b)
			if a == 98 {
				break
			}
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
	}
	ft.PrintRune('\n')
}
