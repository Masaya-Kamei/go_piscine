package main

import (
	"fmt"
	"piscine"
)

func printBTree(root *TreeNode) {
	if root == nil {
		return
	}
	for {

	}
}

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	fmt.Println(root.Left.Data)
	fmt.Println(root.Data)
	fmt.Println(root.Right.Left.Data)
	fmt.Println(root.Right.Data)
}
