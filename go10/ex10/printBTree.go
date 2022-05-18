package main

import (
	"fmt"
	"piscine"
)

type TreeNode = piscine.TreeNode

func printTreeNodeLevel(root *TreeNode, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}
	if root == nil {
		fmt.Println("nil")
		return
	} else {
		fmt.Println(root.Data)
	}
	if root.Right != nil || root.Left != nil {
		printTreeNodeLevel(root.Right, depth+1)
		printTreeNodeLevel(root.Left, depth+1)
	}
}

func printBTree(root *TreeNode) {
	printTreeNodeLevel(root, 0)
}
