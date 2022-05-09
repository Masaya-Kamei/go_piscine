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
	piscine.ListPushBack(link, "Hello")
	printLink(link)
	piscine.ListPushBack(link, "there")
	printLink(link)
	piscine.ListPushBack(link, "how are you")
	printLink(link)

	ft.PrintRune('\n')
	piscine.ListPushBack(link, 42)
	printLink(link)
	piscine.ListPushBack(link, nil)
	printLink(link)
	piscine.ListPushBack(nil, "how are you")
}
