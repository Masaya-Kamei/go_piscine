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
	printBTree(root)
	ft.PrintRune('\n')

	fmt.Println(piscine.BTreeIsBinary(root))
	root.Left.Left = &TreeNode{Data: "8", Parent: root.Left}
	fmt.Println(piscine.BTreeIsBinary(root))
	fmt.Println(piscine.BTreeIsBinary(&TreeNode{Data: "4"}))
	fmt.Println(piscine.BTreeIsBinary(nil))
}
