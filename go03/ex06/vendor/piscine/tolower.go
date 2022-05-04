package piscine

func ToLower(s string) string {
	rs := []rune(s)
	for i := range rs {
		if rs[i] >= 'A' && rs[i] <= 'Z' {
			rs[i] += 32
		}
	}
	return string(rs)
}
