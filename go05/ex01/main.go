package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.MakeRange(5, 10))
	fmt.Println(piscine.MakeRange(10, 5))

	fmt.Println()
	fmt.Println(piscine.MakeRange(5, 5))
	fmt.Println(piscine.MakeRange(5, 6))
	fmt.Println(piscine.MakeRange(9223372036854775806, 9223372036854775807))
	fmt.Println(piscine.MakeRange(9223372036854775807, 9223372036854775807))
	fmt.Println(piscine.MakeRange(-9223372036854775808, -9223372036854775807))
	// intMin, intMax := -9223372036854775808, 9223372036854775807
	// fmt.Println(piscine.MakeRange(0, intMax))
	// fmt.Println(piscine.MakeRange(-1, intMax))
	// fmt.Println(piscine.MakeRange(intMin, intMax))
}
