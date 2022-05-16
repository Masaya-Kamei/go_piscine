package main

import (
	"fmt"
	"piscine"
)

type List = piscine.List
type NodeL = piscine.NodeL

func printList(l *List) {
	fmt.Print("[")
	for link := l.Head; link != nil; link = link.Next {
		fmt.Print(link.Data, " -> ")
	}
	fmt.Print("nil", "]")
	gD := func(l *NodeL) interface{} {
		if l == nil {
			return "nil"
		} else {
			return l.Data
		}
	}
	fmt.Printf("  Head:%v, Tail:%v\n", gD(l.Head), gD(l.Tail))
}
