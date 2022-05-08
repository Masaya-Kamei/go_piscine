package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)

	fmt.Println()
	piscine.PrintWordsTables([]string{"a", "", "b"})
	piscine.PrintWordsTables(nil)
}
