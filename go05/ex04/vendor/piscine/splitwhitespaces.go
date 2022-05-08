package piscine

func getSpacesIndex(s string) int {
	for i, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			return i
		}
	}
	return -1
}

func SplitWhiteSpaces(s string) []string {
	strs := []string{}
	for s != "" {
		spacesIndex := getSpacesIndex(s)
		if spacesIndex == -1 {
			break
		}
		strs = append(strs, s[:spacesIndex])
		s = s[spacesIndex+1:]
	}
	strs = append(strs, s)
	return strs
}
