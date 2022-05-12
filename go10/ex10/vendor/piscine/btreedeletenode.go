package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func isPartOfBTree(root, node *TreeNode) bool {
	for n := node; n != nil; n = n.Parent {
		if n == root {
			return true
		}
	}
	return false
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil || !isPartOfBTree(root, node) {
		return root
	}
	if node.Left == nil && node.Right == nil {
		if node == root {
			return nil
		}
		BTreeTransplant(node.Parent, node, nil)
	} else if node.Left != nil {
		max := BTreeMax(node.Left)
		BTreeTransplant(max.Parent, max, max.Left)
		node.Data = max.Data
	} else if node.Right != nil {
		min := BTreeMin(node.Right)
		BTreeTransplant(min.Parent, min, min.Right)
		node.Data = min.Data
	}
	return root
}
