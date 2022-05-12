package main

import (
	"fmt"
	"ft"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	fmt.Println("Initial root")
	printBTree(root)
	ft.PrintRune('\n')

	fmt.Println("SearchItem (root, 7)")
	printBTree(piscine.BTreeSearchItem(root, "7"))
	ft.PrintRune('\n')

	fmt.Println("SearchItem (root, 1)")
	printBTree(piscine.BTreeSearchItem(root, "1"))
	ft.PrintRune('\n')

	fmt.Println("SearchItem (&piscine.TreeNode{Data: \"4\"}, 4)")
	printBTree(piscine.BTreeSearchItem(&piscine.TreeNode{Data: "4"}, "4"))
	ft.PrintRune('\n')

	fmt.Println("SearchItem (root, 0)")
	printBTree(piscine.BTreeSearchItem(root, "0"))
	ft.PrintRune('\n')

	fmt.Println("SearchItem (nil, 7)")
	printBTree(piscine.BTreeSearchItem(nil, "7"))
}
