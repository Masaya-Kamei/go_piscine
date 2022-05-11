package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	n := root
	for n.Right != nil {
		n = n.Right
	}
	return n
}
