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

	fmt.Println("ApplyInorder (root, fmt.Println)")
	piscine.BTreeApplyInorder(root, fmt.Println)

	fmt.Println("ApplyInorder (&piscine.TreeNode{Data: \"4\"}, fmt.Println)")
	piscine.BTreeApplyInorder(&piscine.TreeNode{Data: "4"}, fmt.Println)

	fmt.Println("ApplyInorder (&piscine.TreeNode{Data: \"4\"}, nil)")
	piscine.BTreeApplyInorder(&piscine.TreeNode{Data: "4"}, nil)

	fmt.Println("ApplyInorder (nil, fmt.Println)")
	piscine.BTreeApplyInorder(nil, fmt.Println)
}
