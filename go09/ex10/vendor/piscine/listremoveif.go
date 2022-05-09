package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func listRemove(l *List, prev *NodeL, target *NodeL) {
	if prev == nil {
		l.Head = target.Next
	} else {
		prev.Next = target.Next
	}
	if target == l.Tail {
		l.Tail = prev
	}
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}
	var prev *NodeL = nil
	for node := l.Head; node != nil; node = node.Next {
		if node.Data == data_ref {
			listRemove(l, prev, node)
		} else {
			prev = node
		}
	}
}
