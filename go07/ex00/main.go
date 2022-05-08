package main

import (
	"ft"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, a)
	ft.PrintRune('\n')

	ft.PrintRune('\n')
	a = []int{0, -9223372036854775808, 9223372036854775807}
	piscine.ForEach(piscine.PrintNbr, a)
	ft.PrintRune('\n')
	piscine.ForEach(piscine.PrintNbr, nil)
	piscine.ForEach(nil, a)
}
