package piscine

func ToUpper(s string) string {
	rs := []rune(s)
	for i := range rs {
		if rs[i] >= 'a' && rs[i] <= 'z' {
			rs[i] -= 32
		}
	}
	return string(rs)
}
