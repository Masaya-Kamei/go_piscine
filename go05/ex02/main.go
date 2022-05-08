package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))

	fmt.Println()
	fmt.Println(piscine.ConcatParams([]string{"a", "", "b"}))
	fmt.Println(piscine.ConcatParams(nil))
}
