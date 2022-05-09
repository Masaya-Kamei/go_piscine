package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFrontNode(l *List, newNode *NodeL) {
	if l == nil {
		return
	}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}

func ListReverse(l *List) {
	if l == nil {
		return
	}
	var next *NodeL
	revl := &List{}
	for node := l.Head; node != nil; node = next {
		next = node.Next
		node.Next = nil
		ListPushFrontNode(revl, node)
	}
	l.Head = revl.Head
	l.Tail = revl.Tail
}
