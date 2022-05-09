package main

import (
	"fmt"
	"ft"
	"piscine"
)

type List = piscine.List
type Node = piscine.NodeL

func printList(l *List) {
	if l == nil {
		return
	}
	link := l.Head
	fmt.Print("[")
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Print(nil, "]")
	if l.Head == nil && l.Tail == nil {
		fmt.Println("  Head:", nil, "Tail:", nil)
	} else {
		fmt.Println("  Head:", l.Head.Data, "Tail:", l.Tail.Data)
	}
}

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
