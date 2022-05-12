package piscine

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
