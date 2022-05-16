package main

import (
	"ft"
	"piscine"
)

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
}
