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

	fmt.Println("Transplant (root, 1, 3)")
	node := p.BTreeSearchItem(root, "1")
	replacement := &p.TreeNode{Data: "3"}
	printBTree(p.BTreeTransplant(root, node, replacement))
	p.BTreeApplyInorder(root, fmt.Println)
	ft.PrintRune('\n')

	fmt.Println("Transplant (root, 3, nil)")
	printBTree(p.BTreeTransplant(root, p.BTreeSearchItem(root, "3"), nil))
	ft.PrintRune('\n')

	fmt.Println("Transplant (root, root, 7)")
	printBTree(p.BTreeTransplant(root, root, p.BTreeSearchItem(root, "7")))
	ft.PrintRune('\n')

	fmt.Println("Transplant (root, root, nil)")
	printBTree(p.BTreeTransplant(root, root, nil))
	ft.PrintRune('\n')

	fmt.Println("Transplant (root, &p.TreeNode{Data: \"5\"}, nil)")
	printBTree(p.BTreeTransplant(root, &p.TreeNode{Data: "5"}, nil))
	ft.PrintRune('\n')

	fmt.Println("Transplant (root, nil, &p.TreeNode{Data: \"6\"})")
	printBTree(p.BTreeTransplant(root, nil, &p.TreeNode{Data: "6"}))
	ft.PrintRune('\n')

	fmt.Println("Transplant (nil, 5, nil)")
	printBTree(p.BTreeTransplant(nil, p.BTreeSearchItem(root, "5"), nil))
}
