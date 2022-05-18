package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	link := &piscine.List{}
	link2 := &piscine.List{}
	piscine.ListPushBack(link, "three")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "1")
	fmt.Println(piscine.ListLast(link))
	fmt.Println(piscine.ListLast(link2))

	ft.PrintRune('\n')
	fmt.Println(piscine.ListLast(nil))
}