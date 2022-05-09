package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(a)
	fmt.Println(unmatch)

	ft.PrintRune('\n')
	fmt.Println(piscine.Unmatch([]int{1, 2, 3, 1, 2, 3, 1}))
	fmt.Println(piscine.Unmatch([]int{42}))
	fmt.Println(piscine.Unmatch([]int{}))
	fmt.Println(piscine.Unmatch(nil))
}
