package piscine

func itoaBase(nbr int, base string) string {
	baseLen := strLen(base)
	indexMap := createIndexMap(base)
	if !isValidBase(base, baseLen, indexMap) {
		return "NV"
	}
	s := ""
	rbase := []rune(base)
	for {
		s = string(rbase[nbr%baseLen]) + s
		nbr /= baseLen
		if nbr == 0 {
			break
		}
	}
	return s
}
