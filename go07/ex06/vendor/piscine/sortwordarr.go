package piscine

func SortWordArr(a []string) {
	for {
		i := 0
		for range a {
			if i >= 1 && a[i-1] > a[i] {
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
