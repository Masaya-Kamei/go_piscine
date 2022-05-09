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
