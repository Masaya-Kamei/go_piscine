package main

import "piscine"

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)

	// a = []string{"", "a", ""}
	// piscine.PrintWordsTables(a)
	// piscine.PrintWordsTables(nil)
}
