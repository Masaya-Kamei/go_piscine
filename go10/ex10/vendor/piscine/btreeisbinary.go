package piscine

func bTreeLessCheck(node *TreeNode, data string) bool {
	if node == nil {
		return true
	}
	ok := node.Data <= data
	return ok && bTreeLessCheck(node.Left, data) && bTreeLessCheck(node.Right, data)
}

func bTreeMoreCheck(node *TreeNode, data string) bool {
	if node == nil {
		return true
	}
	ok := node.Data > data
	return ok && bTreeMoreCheck(node.Left, data) && bTreeMoreCheck(node.Right, data)
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ok := bTreeLessCheck(root.Left, root.Data) && bTreeMoreCheck(root.Right, root.Data)
	return ok && BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
