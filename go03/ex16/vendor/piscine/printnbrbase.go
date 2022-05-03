package piscine

import "ft"

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func abs(nbr int) int {
	if nbr < 0 {
		return -nbr
	}
	return nbr
}

func isValidBase(base string, baseLen int) bool {
	if baseLen <= 1 {
		return false
	}
	m := map[rune]bool{}
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = true
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	baseLen := strLen(base)
	if !isValidBase(base, baseLen) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	if nbr < 0 {
		ft.PrintRune('-')
	}
	rbase := []rune(base)
	for {
		defer ft.PrintRune(rbase[abs(nbr%baseLen)])
		nbr /= baseLen
		if nbr == 0 {
			break
		}
	}
}
