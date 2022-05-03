package piscine

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func NRune(s string, n int) rune {
	if n < 1 || n > strLen(s) {
		return 0
	}
	return []rune(s)[n-1]
}
