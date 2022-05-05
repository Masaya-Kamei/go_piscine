package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))

	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("  Hello how are you?  "))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello  how  \nare \tyou?"))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" あい  うえ "))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("    \t\n"))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(""))
}
