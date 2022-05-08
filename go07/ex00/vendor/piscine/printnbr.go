package piscine

import "ft"

func abs(nbr int) int {
	if nbr < 0 {
		return -nbr
	}
	return nbr
}

func PrintNbr(nbr int) {
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
