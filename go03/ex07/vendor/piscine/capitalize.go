package piscine

func isAlNum(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

func toUpperRune(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func toLowerRune(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func Capitalize(s string) string {
	rs := []rune(s)
	for i := range rs {
		if i == 0 || !isAlNum(rs[i-1]) {
			rs[i] = toUpperRune(rs[i])
		} else {
			rs[i] = toLowerRune(rs[i])
		}
	}
	return string(rs)
}
