package main

import (
	"fmt"
	"piscine"
)

func PrintElem(node *NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *NodeL) {
	node.Data = 2
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

	piscine.ListForEachIf(&piscine.List{}, PrintElem, piscine.IsPositiveNode)
	piscine.ListForEachIf(nil, PrintElem, piscine.IsPositiveNode)
	piscine.ListForEachIf(&piscine.List{}, nil, piscine.IsPositiveNode)
	piscine.ListForEachIf(&piscine.List{}, PrintElem, nil)
}
