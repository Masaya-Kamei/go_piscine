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
	link := &piscine.List{}
	printList(link)
	piscine.ListPushFront(link, "Hello")
	printList(link)
	piscine.ListPushFront(link, "there")
	printList(link)
	piscine.ListPushFront(link, "how are you")
	printList(link)

	ft.PrintRune('\n')
	piscine.ListPushFront(link, 42)
	printList(link)
	piscine.ListPushFront(link, nil)
	printList(link)
	piscine.ListPushFront(nil, 42)
}
