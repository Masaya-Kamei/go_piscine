package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	median := piscine.Median(2, 3, 8, 5, 7)
	fmt.Println(median)

	ft.PrintRune('\n')
	fmt.Println(piscine.Median(4, 1, 3, 5, 2))
	fmt.Println(piscine.Median(0, 0, 0, 0, 0))
}
