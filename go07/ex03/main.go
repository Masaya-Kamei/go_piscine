package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)

	ft.PrintRune('\n')
	fmt.Println(piscine.CountIf(piscine.IsNumeric, []string{"1", "2", "3"}))
	fmt.Println(piscine.CountIf(piscine.IsNumeric, nil))
	fmt.Println(piscine.CountIf(nil, []string{"4"}))
}
