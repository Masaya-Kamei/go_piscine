package main

import (
	"fmt"
	"ft"
	"piscine"
)

const (
	intMax = int(^uint(0) >> 1)
	intMin = -intMax - 1
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)

	ft.PrintRune('\n')
	fmt.Println(piscine.Max([]int{intMin, intMin}))
	fmt.Println(piscine.Max([]int{intMax}))
	fmt.Println(piscine.Max([]int{}))
	fmt.Println(piscine.Max(nil))
}
