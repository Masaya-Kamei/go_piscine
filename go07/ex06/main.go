package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.SortWordArr(result)
	fmt.Println(result)

	ft.PrintRune('\n')
	result = []string{"abc2", "abc3", "abc1"}
	piscine.SortWordArr(result)
	fmt.Println(result)
	result = []string{"abc2", "", "abc1"}
	piscine.SortWordArr(result)
	fmt.Println(result)
	result = nil
	piscine.SortWordArr(result)
	fmt.Println(result)
}
