package main

import (
	"fmt"
	"ft"
	"piscine"
)

type List = piscine.List

func printList(l *List) {
	if l == nil {
		return
	}
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil, " Head:", l.Head, " Tail:", l.Tail)
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
