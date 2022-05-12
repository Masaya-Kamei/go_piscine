package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	n := root
	for n.Left != nil {
		n = n.Left
	}
	return n
}
