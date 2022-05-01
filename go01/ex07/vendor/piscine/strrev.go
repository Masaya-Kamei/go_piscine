package piscine

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func StrRev(s string) string {
	r := []rune(s)
	i, j := 0, strLen(s)-1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
