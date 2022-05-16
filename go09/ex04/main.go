package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	link := &List{}
	piscine.ListPushBack(link, "I")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "something")
	piscine.ListPushBack(link, 2)
	fmt.Println("------list------")
	printList(link)
	piscine.ListClear(link)
	fmt.Println("------updated list------")
	printList(link)

	ft.PrintRune('\n')
	piscine.ListClear(nil)
}
