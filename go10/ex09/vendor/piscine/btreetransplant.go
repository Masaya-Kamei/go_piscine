package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func isPartOfBTree(root, node *TreeNode) bool {
	n := node
	for n.Parent != nil {
		n = n.Parent
	}
	return n == root
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil || !isPartOfBTree(root, node) {
		return root
	}
	switch p := node.Parent; {
	case p == nil:
		root = rplc
	case p.Left == node:
		p.Left = rplc
	case p.Right == node:
		p.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
