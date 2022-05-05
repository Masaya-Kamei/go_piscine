package piscine

func appendToStrs(strs []string, s string, wordStart, wordEnd int) []string {
	switch {
	case wordStart == -1:
		return strs
	case wordEnd == -1:
		return append(strs, s[wordStart:])
	default:
		return append(strs, s[wordStart:wordEnd])
	}
}

func SplitWhiteSpaces(s string) []string {
	var strs []string = nil
	wordStart := -1
	for i, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			strs = appendToStrs(strs, s, wordStart, i)
			wordStart = -1
		} else if wordStart == -1 {
			wordStart = i
		}
	}
	strs = appendToStrs(strs, s, wordStart, -1)
	return strs
}
