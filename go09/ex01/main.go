package main

import (
	"fmt"
	"ft"
	"piscine"
)

func printLink(link *piscine.List) {
	head := link.Head
	fmt.Printf("[")
	for head != nil {
		fmt.Printf("%#v, ", head.Data)
		head = head.Next
	}
	fmt.Printf("] ")
	if link.Head != nil {
		fmt.Printf("Head:%#v ", link.Head.Data)
		fmt.Printf("Tail:%#v", link.Tail.Data)
	}
	fmt.Printf("\n")
}

func main() {
	link := &piscine.List{}
	printLink(link)
	piscine.ListPushFront(link, "Hello")
	printLink(link)
	piscine.ListPushFront(link, "there")
	printLink(link)
	piscine.ListPushFront(link, "how are you")
	printLink(link)

	ft.PrintRune('\n')
	piscine.ListPushFront(link, 42)
	printLink(link)
	piscine.ListPushFront(link, nil)
	printLink(link)
	piscine.ListPushFront(nil, 42)
}
