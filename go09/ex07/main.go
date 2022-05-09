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
