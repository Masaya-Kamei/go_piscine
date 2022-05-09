package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	link := &piscine.List{}
	fmt.Println(piscine.ListSize(link))
	piscine.ListPushFront(link, "Hello")
	fmt.Println(piscine.ListSize(link))
	piscine.ListPushFront(link, "2")
	fmt.Println(piscine.ListSize(link))
	piscine.ListPushFront(link, "you")
	fmt.Println(piscine.ListSize(link))
	piscine.ListPushFront(link, "man")
	fmt.Println(piscine.ListSize(link))

	ft.PrintRune('\n')
	fmt.Println(piscine.ListSize(nil))
}
