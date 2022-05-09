package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))

	ft.PrintRune('\n')
	fmt.Println(piscine.Join([]string{"Hello!"}, ""))
	fmt.Println(piscine.Join([]string{"Hello!", " How", " are", " you?"}, ""))
	fmt.Println(piscine.Join([]string{"", "", ""}, ":"))
	fmt.Println(piscine.Join(nil, ":"))
}
