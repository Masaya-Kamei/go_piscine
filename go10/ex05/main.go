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

	fmt.Println("IsBinary (root)")
	fmt.Println(piscine.BTreeIsBinary(root))

	fmt.Println("IsBinary (root)  [Add 8]")
	root.Left.Right = &TreeNode{Data: "8", Parent: root.Left}
	fmt.Println(piscine.BTreeIsBinary(root))

	fmt.Println("IsBinary (root)  [Add 3]")
	root.Left.Right.Data = "3"
	fmt.Println(piscine.BTreeIsBinary(root))

	fmt.Println("IsBinary (&TreeNode{Data: \"4\"})")
	fmt.Println(piscine.BTreeIsBinary(&TreeNode{Data: "4"}))

	fmt.Println("IsBinary (nil)")
	fmt.Println(piscine.BTreeIsBinary(nil))
}
