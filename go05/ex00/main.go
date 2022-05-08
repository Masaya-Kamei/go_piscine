package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10))
	fmt.Println(piscine.AppendRange(10, 5))

	fmt.Println()
	fmt.Println(piscine.AppendRange(5, 5))
	fmt.Println(piscine.AppendRange(5, 6))
	fmt.Println(piscine.AppendRange(9223372036854775806, 9223372036854775807))
	fmt.Println(piscine.AppendRange(9223372036854775807, 9223372036854775807))
	fmt.Println(piscine.AppendRange(-9223372036854775808, -9223372036854775807))
}
