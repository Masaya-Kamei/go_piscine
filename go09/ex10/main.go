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
	link2 := &piscine.List{}
	fmt.Println("----normal state----")
	piscine.ListPushBack(link2, 1)
	piscine.ListPushBack(link2, 1)
	piscine.ListPushBack(link2, 1)
	printList(link2)
	piscine.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	printList(link2)
	fmt.Println()
	fmt.Println("----normal state----")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "There")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "How")
	piscine.ListPushBack(link, "are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 1)
	printList(link)
	piscine.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	printList(link)

	ft.PrintRune('\n')
	piscine.ListRemoveIf(&piscine.List{}, 1)
	piscine.ListRemoveIf(nil, 1)
}
