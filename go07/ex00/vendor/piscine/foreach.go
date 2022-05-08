package piscine

func ForEach(f func(int), a []int) {
	if f == nil {
		return
	}
	for _, b := range a {
		f(b)
	}
}
