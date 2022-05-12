package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return (root.Left == nil || root.Left.Data < root.Data) &&
		(root.Right == nil || root.Data < root.Right.Data) &&
		BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
