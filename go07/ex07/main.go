package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.AdvancedSortWordArr(result, piscine.Compare)
	fmt.Println(result)

	ft.PrintRune('\n')
	result = []string{"abc3", "abc2", "abc1"}
	piscine.AdvancedSortWordArr(result, piscine.Compare)
	fmt.Println(result)
	result = []string{"abc2", "", "abc1"}
	piscine.AdvancedSortWordArr(result, piscine.Compare)
	fmt.Println(result)
	piscine.AdvancedSortWordArr(nil, piscine.Compare)
	piscine.AdvancedSortWordArr([]string{"a"}, nil)
}
