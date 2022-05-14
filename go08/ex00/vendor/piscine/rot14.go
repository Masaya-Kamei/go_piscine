package piscine

func Rot14(s string) string {
	rs := []rune(s)
	for i, r := range rs {
		if r >= 'A' && r <= 'Z' {
			rs[i] = (r-'A'+14)%26 + 'A'
		} else if r >= 'a' && r <= 'z' {
			rs[i] = (r-'a'+14)%26 + 'a'
		}
	}
	return string(rs)
}
