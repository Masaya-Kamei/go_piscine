package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if f == nil {
		return false
	}
	for i := range a {
		if i >= 1 && f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}
