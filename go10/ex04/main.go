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

	fmt.Println("LevelCount (root)")
	fmt.Println(piscine.BTreeLevelCount(root))

	fmt.Println("LevelCount (InsertData(root, \"6\"))")
	fmt.Println(piscine.BTreeLevelCount(piscine.BTreeInsertData(root, "6")))

	fmt.Println("LevelCount (&piscine.TreeNode{Data: \"4\"})")
	fmt.Println(piscine.BTreeLevelCount(&piscine.TreeNode{Data: "4"}))

	fmt.Println("LevelCount (nil)")
	fmt.Println(piscine.BTreeLevelCount(nil))
}
