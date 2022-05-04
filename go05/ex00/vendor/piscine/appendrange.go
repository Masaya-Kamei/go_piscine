package piscine

func AppendRange(min, max int) []int {
	var nbrs []int = nil
	for i := min; i < max; i++ {
		nbrs = append(nbrs, i)
	}
	return nbrs
}
