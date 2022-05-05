package piscine

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

func itoaBase(nbr int, base string) string {
	baseLen := strLen(base)
	if !isValidBase(base, baseLen) {
		return "NV"
	}
	s := ""
	if nbr < 0 {
		s += "-"
	}
	rbase := []rune(base)
	for {
		s = string(rbase[abs(nbr%baseLen)]) + s
		nbr /= baseLen
		if nbr == 0 {
			break
		}
	}
	return s
}
