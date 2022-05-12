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
	printBTree(root)
	ft.PrintRune('\n')

	node := p.BTreeSearchItem(root, "1")
	replacement := &p.TreeNode{Data: "3"}
	printBTree(p.BTreeTransplant(root, node, replacement))
	p.BTreeApplyInorder(root, fmt.Println)
	ft.PrintRune('\n')

	printBTree(p.BTreeTransplant(root, p.BTreeSearchItem(root, "3"), nil))
	ft.PrintRune('\n')
	printBTree(p.BTreeTransplant(root, root, p.BTreeSearchItem(root, "7")))
	ft.PrintRune('\n')
	printBTree(p.BTreeTransplant(root, root, nil))

	fmt.Println("Invalid\n")
	printBTree(p.BTreeTransplant(root, &p.TreeNode{Data: "5"}, nil))
	ft.PrintRune('\n')
	printBTree(p.BTreeTransplant(root, nil, &p.TreeNode{Data: "6"}))
	printBTree(p.BTreeTransplant(nil, p.BTreeSearchItem(root, "5"), nil))
}
