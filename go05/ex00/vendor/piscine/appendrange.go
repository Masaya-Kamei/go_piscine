package piscine

func AppendRange(min, max int) []int {
	var nbrs []int
	for i := min; i < max; i++ {
		nbrs = append(nbrs, i)
	}
	return nbrs
}
