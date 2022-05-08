package piscine

func Any(f func(string) bool, a []string) bool {
	if f == nil {
		return false
	}
	for _, aa := range a {
		if f(aa) {
			return true
		}
	}
	return false
}
