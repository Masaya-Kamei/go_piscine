package main

import (
	"fmt"
	"ft"
	"piscine"
)

type List = piscine.List

func printList(l *List) {
	fmt.Print("[")
	for link := l.Head; link != nil; link = link.Next {
		fmt.Print(link.Data, " -> ")
	}
	fmt.Print(nil, "]")
	if l.Head == nil && l.Tail == nil {
		fmt.Println("  Head:", nil, "Tail:", nil)
	} else {
		fmt.Println("  Head:", l.Head.Data, "Tail:", l.Tail.Data)
	}
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
