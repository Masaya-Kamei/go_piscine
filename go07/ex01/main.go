package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, a)
	fmt.Println(result)

	ft.PrintRune('\n')
	fmt.Println(piscine.Map(piscine.IsPrime, []int{-1, 0, 1, 2}))
	fmt.Println(piscine.Map(piscine.IsPrime, nil))
}
