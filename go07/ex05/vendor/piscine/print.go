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

func abs(nbr int) int {
	if nbr < 0 {
		return -nbr
	}
	return nbr
}

func printNbr(nbr int) {
	if nbr < 0 {
		ft.PrintRune('-')
	}
	for {
		defer ft.PrintRune(rune(abs(nbr%10) + '0'))
		nbr /= 10
		if nbr == 0 {
			break
		}
	}
}

func printNbrEndl(nbr int) {
	printNbr(nbr)
	ft.PrintRune('\n')
}
