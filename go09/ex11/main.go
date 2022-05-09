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
	piscine.ListPushBack(link, "a")
	piscine.ListPushBack(link, "b")
	piscine.ListPushBack(link, "c")
	piscine.ListPushBack(link, "d")
	fmt.Println("-----first List------")
	printList(link)
	piscine.ListPushBack(link2, "e")
	piscine.ListPushBack(link2, "f")
	piscine.ListPushBack(link2, "g")
	piscine.ListPushBack(link2, "h")
	fmt.Println("-----second List------")
	printList(link2)
	fmt.Println("-----Merged List-----")
	piscine.ListMerge(link, link2)
	printList(link)

	ft.PrintRune('\n')

	piscine.ListMerge(link, &piscine.List{})
	printList(link)

	link3 := &piscine.List{}
	piscine.ListMerge(link3, &piscine.List{})
	printList(link3)

	piscine.ListMerge(link, nil)
	piscine.ListMerge(nil, link)
}
