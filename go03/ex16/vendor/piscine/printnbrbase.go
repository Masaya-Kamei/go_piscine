package piscine

import "ft"

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
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
		printStr("NV")
		return
	}
	if nbr < 0 {
		ft.PrintRune('-')
	}
	rbase := []rune(base)
	s := ""
	for {
		s = string(rbase[abs(nbr%baseLen)]) + s
		nbr /= baseLen
		if nbr == 0 {
			break
		}
	}
	printStr(s)
}
