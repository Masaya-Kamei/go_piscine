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

	fmt.Println("Insert (root, 0)")
	printBTree(piscine.BTreeInsertData(root, "0"))
	ft.PrintRune('\n')

	fmt.Println("Insert (root, 1)")
	printBTree(piscine.BTreeInsertData(root, "1"))
	ft.PrintRune('\n')

	fmt.Println("Insert (nil, 1)")
	printBTree(piscine.BTreeInsertData(nil, "1"))
}
