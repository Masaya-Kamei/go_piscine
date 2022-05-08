package piscine

func bytesLen(bytes []byte) int {
	i := 0
	for range bytes {
		i++
	}
	return i
}

func getSepIndex(s string, toFind string) int {
	sByteLen := bytesLen([]byte(s))
	toFindByteLen := bytesLen([]byte(toFind))
	for i := 0; i+toFindByteLen <= sByteLen; i++ {
		if s[i:i+toFindByteLen] == toFind {
			return i
		}
	}
	return -1
}

func splitEachRune(s string) []string {
	strs := []string{}
	for _, r := range s {
		strs = append(strs, string(r))
	}
	return strs
}

func Split(s string, sep string) []string {
	if sep == "" {
		return splitEachRune(s)
	}

	strs := []string{}
	sepByteLen := bytesLen([]byte(sep))
	for s != "" {
		sepIndex := getSepIndex(s, sep)
		if sepIndex == -1 {
			break
		}
		strs = append(strs, s[:sepIndex])
		s = s[sepIndex+sepByteLen:]
	}
	strs = append(strs, s)
	return strs
}
