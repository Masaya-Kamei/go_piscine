package piscine

import (
	"ft"
	"os"
)

func printStrEndl(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func sortStringTable(table []string) {
	if table == nil {
		return
	}
	i := 0
	for range table {
		if i >= 1 && table[i-1] > table[i] {
			table[i-1], table[i] = table[i], table[i-1]
		}
		i++
	}
	if i >= 3 {
		sortStringTable(table[:i-1])
	}
}

func SortParams() {
	strs := os.Args[1:]
	sortStringTable(strs)
	for _, s := range strs {
		printStrEndl(s)
	}
}
