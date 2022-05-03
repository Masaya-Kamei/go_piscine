package piscine

func isAlNum(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func IsAlpha(s string) bool {
	for _, r := range s {
		if !isAlNum(r) {
			return false
		}
	}
	return true
}
