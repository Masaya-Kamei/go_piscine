package main

import (
	"fmt"
	"ft"
	"piscine"
)

func f(a, b int) int {
	switch {
	case a > b:
		return 1
	case a == b:
		return 0
	default:
		return -1
	}
}

func fReverce(a, b int) int {
	switch {
	case a < b:
		return 1
	case a == b:
		return 0
	default:
		return -1
	}
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)
	fmt.Println(result1)
	fmt.Println(result2)

	ft.PrintRune('\n')
	fmt.Println(piscine.IsSorted(f, []int{0, 1, 1, 1}))
	fmt.Println(piscine.IsSorted(f, []int{0}))
	fmt.Println(piscine.IsSorted(f, []int{}))
	fmt.Println(piscine.IsSorted(f, nil))
	fmt.Println(piscine.IsSorted(nil, []int{1, 0}))
	fmt.Println(piscine.IsSorted(fReverce, []int{0, 1, 2}))
	fmt.Println(piscine.IsSorted(fReverce, []int{2, 1, 0}))
}
