package piscine

func ToLower(s string) string {
	rs := []rune(s)
	for i := range s {
		if rs[i] >= 'A' && rs[i] <= 'Z' {
			rs[i] += 32
		}
	}
	return string(rs)
}
