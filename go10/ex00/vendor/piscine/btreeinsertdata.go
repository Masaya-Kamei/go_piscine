package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
			return root
		}
		return BTreeInsertData(root.Left, data).Parent
	} else if data > root.Data {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
			return root
		}
		return BTreeInsertData(root.Right, data).Parent
	}
	return root
}
