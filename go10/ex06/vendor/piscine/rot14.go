package piscine

func Rot14(s string) string {
	alphabetLarge := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetSmall := "abcdefghijklmnopqrstuvwxyz"
	rs := []rune(s)
	for i, r := range rs {
		if r >= 'A' && r <= 'Z' {
			rs[i] = rune(alphabetLarge[(r-'A'+14)%26])
		} else if r >= 'a' && r <= 'z' {
			rs[i] = rune(alphabetSmall[(r-'a'+14)%26])
		}
	}
	return string(rs)
}
