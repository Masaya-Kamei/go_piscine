package main

import (
	"fmt"
	"piscine"
)

type TreeNode = piscine.TreeNode

func printTreeNode(n *TreeNode) {
	gD := func(n *TreeNode) string {
		if n == nil {
			return " "
		}
		return n.Data
	}
	fmt.Printf("(%v->)%v->[%v, %v] ", gD(n.Parent), n.Data, gD(n.Left), gD(n.Right))
}

func printTreeNodeLevel(root *TreeNode, depth int) bool {
	if root == nil {
		return false
	}
	if depth == 0 {
		printTreeNode(root)
		return true
	}
	ok1 := printTreeNodeLevel(root.Left, depth-1)
	ok2 := printTreeNodeLevel(root.Right, depth-1)
	return ok1 || ok2
}

func printBTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	ok := true
	for depth := 0; ok; depth++ {
		if depth != 0 {
			fmt.Println()
		}
		ok = printTreeNodeLevel(root, depth)
	}
}
