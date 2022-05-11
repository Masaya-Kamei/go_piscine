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

	fmt.Println(piscine.BTreeLevelCount(root))
	fmt.Println(piscine.BTreeLevelCount(piscine.BTreeInsertData(root, "6")))
	fmt.Println(piscine.BTreeLevelCount(&piscine.TreeNode{Data: "4"}))
	fmt.Println(piscine.BTreeLevelCount(nil))
}
