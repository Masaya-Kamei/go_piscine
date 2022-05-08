package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	if f == nil {
		return
	}
	for {
		i := 0
		for range a {
			if i >= 1 && f(a[i-1], a[i]) > 0 {
				a[i-1], a[i] = a[i], a[i-1]
			}
			i++
		}
		if i <= 2 {
			break
		}
		a = a[:i-1]
	}
}
