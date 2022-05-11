package main

import (
	"ft"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	printBTree(root)
	ft.PrintRune('\n')

	printBTree(piscine.BTreeInsertData(root, "1"))
	ft.PrintRune('\n')

	printBTree(piscine.BTreeInsertData(nil, "1"))
}
