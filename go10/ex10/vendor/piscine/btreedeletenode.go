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

func getAltNode(node *TreeNode) *TreeNode {
	var altNode *TreeNode

	if node.Left == nil && node.Right == nil {
		altNode = nil
	} else if node.Left != nil {
		altNode = BTreeMax(node.Left)
		BTreeTransplant(altNode, altNode, altNode.Left)
	} else if node.Right != nil {
		altNode = BTreeMin(node.Right)
		BTreeTransplant(altNode, altNode, altNode.Right)
	}
	return altNode
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil || !isPartOfBTree(root, node) {
		return root
	}

	altNode := getAltNode(node)
	if altNode != nil {
		altNode.Left = node.Left
		altNode.Right = node.Right
		if altNode.Left != nil {
			altNode.Left.Parent = altNode
		}
		if altNode.Right != nil {
			altNode.Right.Parent = altNode
		}
	}
	root = BTreeTransplant(root, node, altNode)
	return root
}
