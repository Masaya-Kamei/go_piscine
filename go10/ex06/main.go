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

	piscine.BTreeApplyByLevel(root, fmt.Println)
	ft.PrintRune('\n')
	piscine.BTreeInsertData(root, "0")
	piscine.BTreeApplyByLevel(root, fmt.Println)
	ft.PrintRune('\n')
	piscine.BTreeApplyByLevel(&piscine.TreeNode{Data: "4"}, fmt.Println)
	piscine.BTreeApplyByLevel(nil, fmt.Println)
	piscine.BTreeApplyByLevel(&piscine.TreeNode{Data: "4"}, nil)
}
