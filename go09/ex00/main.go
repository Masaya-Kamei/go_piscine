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
	printList(link)
	piscine.ListPushBack(link, "Hello")
	printList(link)
	piscine.ListPushBack(link, "there")
	printList(link)
	piscine.ListPushBack(link, "how are you")
	printList(link)

	ft.PrintRune('\n')
	piscine.ListPushBack(link, 42)
	printList(link)
	piscine.ListPushBack(link, nil)
	printList(link)
	piscine.ListPushBack(nil, "how are you")
}
