package piscine

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func Index(s string, toFind string) int {
	if s == "" || toFind == "" {
		return -1
	}
	sLen := strLen(s)
	toFindLen := strLen(toFind)
	rs := []rune(s)
	for i := 0; i+toFindLen <= sLen; i++ {
		if string(rs[i:i+toFindLen]) == toFind {
			return i
		}
	}
	return -1
}
