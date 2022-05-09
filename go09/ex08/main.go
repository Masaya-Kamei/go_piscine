package main

import (
	"fmt"
	"ft"
	"piscine"
)

type List = piscine.List
type NodeL = piscine.NodeL

func PrintElem(node *NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *NodeL) {
	node.Data = 2
}

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
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "there")
	piscine.ListPushBack(link, 23)
	piscine.ListPushBack(link, "!")
	piscine.ListPushBack(link, 54)
	printList(link)
	piscine.ListForEachIf(link, PrintElem, piscine.IsPositiveNode)
	piscine.ListForEachIf(link, StringToInt, piscine.IsAlNode)
	printList(link)

	ft.PrintRune('\n')
	piscine.ListForEachIf(&piscine.List{}, PrintElem, piscine.IsPositiveNode)
	piscine.ListForEachIf(nil, PrintElem, piscine.IsPositiveNode)
	piscine.ListForEachIf(&piscine.List{}, nil, piscine.IsPositiveNode)
	piscine.ListForEachIf(&piscine.List{}, PrintElem, nil)
	printList(link)
}
