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

	fmt.Println("Min (root)")
	printBTree(piscine.BTreeMin(root))
	ft.PrintRune('\n')

	fmt.Println("Min (&piscine.TreeNode{Data: \"4\"})")
	printBTree(piscine.BTreeMin(&piscine.TreeNode{Data: "4"}))
	ft.PrintRune('\n')

	fmt.Println("Min (nil)")
	printBTree(piscine.BTreeMin(nil))
}
