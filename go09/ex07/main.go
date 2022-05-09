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
	piscine.ListPushBack(link, "1")
	piscine.ListPushBack(link, "2")
	piscine.ListPushBack(link, "3")
	piscine.ListPushBack(link, "5")
	printList(link)
	piscine.ListForEach(link, piscine.Add2_node)
	printList(link)

	ft.PrintRune('\n')
	piscine.ListForEach(link, piscine.Subtract3_node)
	printList(link)
	piscine.ListForEach(&piscine.List{}, piscine.Subtract3_node)
	piscine.ListForEach(nil, piscine.Subtract3_node)
	piscine.ListForEach(link, nil)
	printList(link)
}
