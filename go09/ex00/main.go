package main

import (
	"ft"
	"piscine"
)

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
