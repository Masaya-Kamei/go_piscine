package piscine

func Map(f func(int) bool, a []int) []bool {
	if f == nil {
		return []bool{}
	}
	b := []bool{}
	for _, aa := range a {
		b = append(b, f(aa))
	}
	return b
}
