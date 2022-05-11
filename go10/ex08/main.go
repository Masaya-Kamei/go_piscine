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

	printTreeNode(piscine.BTreeMin(root))
	ft.PrintRune('\n')
	printTreeNode(piscine.BTreeMin(&piscine.TreeNode{Data: "4"}))
	ft.PrintRune('\n')
	printTreeNode(piscine.BTreeMin(nil))
}
