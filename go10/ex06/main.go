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

	fmt.Println("ApplyByLevel (root)")
	piscine.BTreeApplyByLevel(root, fmt.Println)

	fmt.Println("ApplyByLevel (Insert (root, 0))")
	piscine.BTreeApplyByLevel(piscine.BTreeInsertData(root, "0"), fmt.Println)

	fmt.Println("ApplyByLevel (&piscine.TreeNode{Data: \"4\"}, fmt.Println)")
	piscine.BTreeApplyByLevel(&piscine.TreeNode{Data: "4"}, fmt.Println)

	fmt.Println("ApplyByLevel (nil, fmt.Println)")
	piscine.BTreeApplyByLevel(nil, fmt.Println)

	fmt.Println("ApplyByLevel (root, nil)")
	piscine.BTreeApplyByLevel(root, nil)
}
