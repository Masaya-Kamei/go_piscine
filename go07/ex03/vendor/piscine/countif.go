package piscine

func CountIf(f func(string) bool, tab []string) int {
	if f == nil {
		return 0
	}
	count := 0
	for _, t := range tab {
		if f(t) {
			count++
		}
	}
	return count
}
