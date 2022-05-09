package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	link := &piscine.List{}
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "hello1")
	piscine.ListPushBack(link, "hello2")
	piscine.ListPushBack(link, "hello3")
	found := piscine.ListFind(link, interface{}("hello2"), piscine.CompStr)
	fmt.Println(found)
	fmt.Println(*found)

	ft.PrintRune('\n')

	found = piscine.ListFind(link, "hello2", piscine.CompStr)
	fmt.Println(found, *found)
	found = piscine.ListFind(link, "hello5", piscine.CompStr)
	fmt.Println(found)

	piscine.ListPushBack(link, 1)
	found = piscine.ListFind(link, 1, piscine.CompStr)
	fmt.Println(found, *found)
	found = piscine.ListFind(link, float64(1), piscine.CompStr)
	fmt.Println(found)

	piscine.ListPushBack(link, nil)
	found = piscine.ListFind(link, nil, piscine.CompStr)
	fmt.Println(found, *found)
	var a *int = nil
	found = piscine.ListFind(link, a, piscine.CompStr)
	fmt.Println(found)

	found = piscine.ListFind(nil, interface{}("hello2"), piscine.CompStr)
	fmt.Println(found)
	found = piscine.ListFind(link, interface{}("hello2"), nil)
	fmt.Println(found)
}
