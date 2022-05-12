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

	fmt.Println("ApplyPreorder (root, fmt.Println)")
	piscine.BTreeApplyPreorder(root, fmt.Println)

	fmt.Println("ApplyPreorder (&piscine.TreeNode{Data: \"4\"}, fmt.Println)")
	piscine.BTreeApplyPreorder(&piscine.TreeNode{Data: "4"}, fmt.Println)

	fmt.Println("ApplyPreorder (&piscine.TreeNode{Data: \"4\"}, nil)")
	piscine.BTreeApplyPreorder(&piscine.TreeNode{Data: "4"}, nil)

	fmt.Println("ApplyPreorder (nil, fmt.Println)")
	piscine.BTreeApplyPreorder(nil, fmt.Println)
}
