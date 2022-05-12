package main

import (
	"fmt"
	"ft"
	p "piscine"
)

func main() {
	root := &p.TreeNode{Data: "4"}
	p.BTreeInsertData(root, "1")
	p.BTreeInsertData(root, "7")
	p.BTreeInsertData(root, "5")
	fmt.Println("Initial root")
	printBTree(root)
	ft.PrintRune('\n')

	fmt.Println("Delete (root, 4)")
	printBTree(p.BTreeDeleteNode(root, p.BTreeSearchItem(root, "4")))
	fmt.Println(p.BTreeIsBinary(root))
	p.BTreeApplyInorder(root, fmt.Println)
	ft.PrintRune('\n')

	fmt.Println("Delete (root, 7)")
	printBTree(p.BTreeDeleteNode(root, p.BTreeSearchItem(root, "7")))
	fmt.Println(p.BTreeIsBinary(root))
	ft.PrintRune('\n')

	fmt.Println("Delete (root, &p.TreeNode{Data: \"8\"})")
	printBTree(p.BTreeDeleteNode(root, &p.TreeNode{Data: "8"}))
	fmt.Println(p.BTreeIsBinary(root))
	ft.PrintRune('\n')

	fmt.Println("Delete (root, 5)")
	printBTree(p.BTreeDeleteNode(root, p.BTreeSearchItem(root, "5")))
	fmt.Println(p.BTreeIsBinary(root))
	ft.PrintRune('\n')

	fmt.Println("Delete (root, 1)")
	root = p.BTreeDeleteNode(root, p.BTreeSearchItem(root, "1"))
	printBTree(root)
	fmt.Println(p.BTreeIsBinary(root))
	ft.PrintRune('\n')

	fmt.Println("Delete (&p.TreeNode{Data: \"4\"}, nil")
	root = p.BTreeDeleteNode(&p.TreeNode{Data: "4"}, nil)
	printBTree(root)
	fmt.Println(p.BTreeIsBinary(root))
	ft.PrintRune('\n')

	fmt.Println("Delete (nil, &p.TreeNode{Data: \"4\"})")
	root = p.BTreeDeleteNode(nil, &p.TreeNode{Data: "4"})
	printBTree(root)
	fmt.Println(p.BTreeIsBinary(root))
}
