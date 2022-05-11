package main

import (
	"fmt"
	"ft"
	"piscine"
)

type TreeNode = piscine.TreeNode

const maxDepth = 4
const maxLength int = 2 * 2 * 2

func storeToTable(tn *TreeNode, table [][maxLength]*TreeNode, depth, i int) {
	if tn == nil || depth == maxDepth {
		return
	}
	table[depth][i] = tn
	storeToTable(tn.Left, table, depth+1, i*2)
	storeToTable(tn.Right, table, depth+1, i*2+1)
}

func gD(n *TreeNode) string {
	if n == nil {
		return " "
	}
	return n.Data
}

func printTreeNode(n *TreeNode) int {
	if n == nil {
		return 0
	}
	fmt.Printf("(%v->)%v->[%v, %v] ", gD(n.Parent), n.Data, gD(n.Left), gD(n.Right))
	return 1
}

func printBTree(root *TreeNode) {
	if root == nil {
		return
	}
	table := [maxDepth][maxLength]*TreeNode{}
	storeToTable(root, table[:], 0, 0)
	for i := 0; i < maxDepth; i++ {
		existFlag := 0
		for j := 0; j < maxLength; j++ {
			existFlag |= printTreeNode(table[i][j])
		}
		if existFlag == 0 {
			break
		}
		ft.PrintRune('\n')
	}
}
