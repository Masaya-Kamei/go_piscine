package piscine

func sortIntegerTable(table []int) {
	for {
		i := 0
		for range table {
			if i >= 1 && table[i-1] > table[i] {
				table[i-1], table[i] = table[i], table[i-1]
			}
			i++
		}
		if i <= 2 {
			break
		}
		table = table[:i-1]
	}
}

func Median(a, b, c, d, e int) int {
	nbrs := []int{a, b, c, d, e}
	sortIntegerTable(nbrs)
	return nbrs[2]
}
