package main

import (
	"ft"
	"piscine"
)

func main() {
	link := &piscine.List{}
	printList(link)
	piscine.ListPushFront(link, "Hello")
	printList(link)
	piscine.ListPushFront(link, "there")
	printList(link)
	piscine.ListPushFront(link, "how are you")
	printList(link)

	ft.PrintRune('\n')
	piscine.ListPushFront(link, 42)
	printList(link)
	piscine.ListPushFront(link, nil)
	printList(link)
	piscine.ListPushFront(nil, 42)
}
