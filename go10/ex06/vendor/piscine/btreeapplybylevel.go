package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func (queue *List) enQueue(data *TreeNode) {
	ListPushBack(queue, data)
}

func (queue *List) deQueue() (*TreeNode, bool) {
	if queue.Head == nil {
		return nil, false
	}
	node := queue.Head.Data.(*TreeNode)
	queue.Head = queue.Head.Next
	if queue.Head == nil {
		queue.Tail = nil
	}
	return node, true
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil || f == nil {
		return
	}
	queue := &List{}
	node := root
	ok := true
	for ok {
		f(node.Data)
		if node.Left != nil {
			queue.enQueue(node.Left)
		}
		if node.Right != nil {
			queue.enQueue(node.Right)
		}
		node, ok = queue.deQueue()
	}
}
