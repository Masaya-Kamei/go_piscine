package main

import (
	"fmt"
	"ft"
	"piscine"
)

type List = piscine.List

func printList(link *List) {
	it := link.Head
	for it != nil {
		fmt.Printf("%v -> ", it.Data)
		it = it.Next
	}
	fmt.Printf("%v", nil)
	fmt.Println(" Head:", link.Head, " Tail:", link.Tail)
}

func main() {
	link := &piscine.List{}
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 2)
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, 4)
	piscine.ListReverse(link)
	printList(link)

	ft.PrintRune('\n')
	link = &piscine.List{}
	piscine.ListReverse(link)
	printList(link)
	piscine.ListPushBack(link, 1)
	piscine.ListReverse(link)
	printList(link)
	piscine.ListPushBack(link, 2)
	piscine.ListReverse(link)
	printList(link)
	piscine.ListReverse(nil)
}
