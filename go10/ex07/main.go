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

	fmt.Println("Max (root)")
	printBTree(piscine.BTreeMax(root))
	ft.PrintRune('\n')

	fmt.Println("Max (&piscine.TreeNode{Data: \"4\"})")
	printBTree(piscine.BTreeMax(&piscine.TreeNode{Data: "4"}))
	ft.PrintRune('\n')

	fmt.Println("Max (nil)")
	printBTree(piscine.BTreeMax(nil))
}
