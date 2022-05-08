package piscine

func getIntSliceLen(a []int) int {
	i := 0
	for range a {
		i++
	}
	return i
}

func Map(f func(int) bool, a []int) []bool {
	b := make([]bool, getIntSliceLen(a))
	for i, aa := range a {
		b[i] = f(aa)
	}
	return b
}
