package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))
	fmt.Println(piscine.AtoiBase("1111101", "01"))
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))

	// fmt.Println(piscine.AtoiBase("一二五", "零一二三四五六七八九"))
	// fmt.Println(piscine.AtoiBase("a", "ab"))
	// fmt.Println(piscine.AtoiBase("", "ab"))
	// fmt.Println(piscine.AtoiBase("125", "01234567890"))
	// fmt.Println(piscine.AtoiBase("000", "0"))
	// fmt.Println(piscine.AtoiBase("125", "+0123456789"))
	// fmt.Println(piscine.AtoiBase("125a", "0123456789"))
}
