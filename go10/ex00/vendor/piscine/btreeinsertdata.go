package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	parent := root
	child := &root
	for *child != nil {
		parent = *child
		if data < parent.Data {
			child = &parent.Left
		} else if data > parent.Data {
			child = &parent.Right
		} else {
			return root
		}
	}
	*child = &TreeNode{Data: data, Parent: parent}
	return root
}
