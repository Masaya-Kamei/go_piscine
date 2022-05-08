package piscine

func MakeRange(min, max int) []int {
	if min >= max || max-min <= 0 {
		return nil
	}
	nbrs := make([]int, max-min)
	for i := 0; i+min < max; i++ {
		nbrs[i] = i + min
	}
	return nbrs
}
