package piscine

func SortIntegerTable(table []int) {
	i := 0
	for range table {
		if i >= 1 && table[i-1] > table[i] {
			table[i-1], table[i] = table[i], table[i-1]
		}
		i++
	}
	if i >= 3 {
		SortIntegerTable(table[:i-1])
	}
}
