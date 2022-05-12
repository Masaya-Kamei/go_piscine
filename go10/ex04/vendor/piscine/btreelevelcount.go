package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func max(nbr1, nbr2 int) int {
	if nbr1 <= nbr2 {
		return nbr2
	} else {
		return nbr1
	}
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(BTreeLevelCount(root.Left), BTreeLevelCount(root.Right))
}
