package piscine

func getIntSliceLen(a []int) int {
	i := 0
	for range a {
		i++
	}
	return i
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if f == nil {
		return false
	}
	len := getIntSliceLen(a)
	for i := 0; i+1 < len; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}
