package main

import (
	"fmt"
	"piscine"
)

func PrintList(l *piscine.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
	n := &piscine.NodeI{Data: data}
	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *piscine.NodeI
	link = piscine.SortListInsert(link, 1)
	PrintList(link)
	link = piscine.SortListInsert(link, 4)
	PrintList(link)
	link = piscine.SortListInsert(link, 9)
	PrintList(link)
	link = piscine.SortListInsert(link, -2)
	PrintList(link)
	link = piscine.SortListInsert(link, 2)
	PrintList(link)
}
