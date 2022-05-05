package piscine

func strLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func appendToStrs(strs []string, rs []rune, wordStart, wordEnd int) []string {
	switch {
	case wordStart == -1:
		return strs
	case wordEnd == -1:
		return append(strs, string(rs[wordStart:]))
	default:
		return append(strs, string(rs[wordStart:wordEnd]))
	}
}

func Split(s, sep string) []string {
	var strs []string = nil
	rs := []rune(s)
	sLen := strLen(s)
	sepLen := strLen(sep)
	wordStart := -1

	for i := 0; i < sLen; i++ {
		if i+sepLen <= sLen && string(rs[i:i+sepLen]) == sep {
			strs = appendToStrs(strs, rs, wordStart, i)
			i += sepLen - 1
			wordStart = -1
		} else if wordStart == -1 {
			wordStart = i
		}
	}
	strs = appendToStrs(strs, rs, wordStart, -1)
	return strs
}
