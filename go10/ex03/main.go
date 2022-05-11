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

	printTreeNode(piscine.BTreeSearchItem(root, "7"))
	ft.PrintRune('\n')
	printTreeNode(piscine.BTreeSearchItem(root, "1"))
	ft.PrintRune('\n')
	printTreeNode(piscine.BTreeSearchItem(&piscine.TreeNode{Data: "4"}, "4"))
	ft.PrintRune('\n')
	printTreeNode(piscine.BTreeSearchItem(root, "0"))
	printTreeNode(piscine.BTreeSearchItem(nil, "7"))
}
