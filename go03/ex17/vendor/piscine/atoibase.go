package piscine

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func createIndexMap(base string) map[rune]int {
	m := map[rune]int{}
	for i, r := range base {
		if _, ok := m[r]; ok {
			return nil
		}
		m[r] = i
	}
	return m
}

func isValidBase(base string, baseLen int) bool {
	if baseLen <= 1 {
		return false
	}
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
	}
	return true
}

func isValidS(s string, indexMap map[rune]int) bool {
	for _, r := range s {
		if _, ok := indexMap[r]; !ok {
			return false
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	baseLen := strLen(base)
	indexMap := createIndexMap(base)
	if indexMap == nil || !isValidBase(base, baseLen) || !isValidS(s, indexMap) {
		return 0
	}
	nbr := 0
	for _, r := range s {
		if nbr*baseLen/baseLen != nbr {
			return 0
		}
		nbr *= baseLen
		if nbr+indexMap[r] < 0 {
			return 0
		}
		nbr += indexMap[r]
	}
	return nbr
}
