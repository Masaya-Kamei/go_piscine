package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

type funcType func(...interface{}) (int, error)

func bTreeApplyLevel(root *TreeNode, f funcType, depth int) bool {
	if root == nil {
		return false
	}
	if depth == 0 {
		f(root.Data)
		return true
	}
	ok1 := bTreeApplyLevel(root.Left, f, depth-1)
	ok2 := bTreeApplyLevel(root.Right, f, depth-1)
	return ok1 || ok2
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil || f == nil {
		return
	}
	ok := true
	for depth := 0; ok; depth++ {
		ok = bTreeApplyLevel(root, f, depth)
	}
}
