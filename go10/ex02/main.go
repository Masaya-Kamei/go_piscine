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

	piscine.BTreeApplyPreorder(root, fmt.Println)
	ft.PrintRune('\n')

	piscine.BTreeApplyPreorder(&piscine.TreeNode{Data: "4"}, fmt.Println)
	piscine.BTreeApplyPreorder(&piscine.TreeNode{Data: "4"}, nil)
	piscine.BTreeApplyPreorder(nil, fmt.Println)
}
