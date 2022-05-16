package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref, Next: nil}
	if l == nil {
		return newNode
	} else if data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}
	for n := l; n != nil; n = n.Next {
		if n.Next == nil || data_ref <= n.Next.Data {
			newNode.Next = n.Next
			n.Next = newNode
			break
		}
	}
	return l
}
