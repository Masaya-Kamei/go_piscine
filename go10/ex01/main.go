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
	piscine.BTreeApplyInorder(root, fmt.Println)
	ft.PrintRune('\n')

	piscine.BTreeApplyInorder(&piscine.TreeNode{Data: "4"}, fmt.Println)
	piscine.BTreeApplyInorder(&piscine.TreeNode{Data: "4"}, nil)
	piscine.BTreeApplyInorder(nil, fmt.Println)
}
