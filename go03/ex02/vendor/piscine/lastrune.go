package piscine

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func LastRune(s string) rune {
	if s == "" {
		return 0
	}
	return []rune(s)[strLen(s)-1]
}
