package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	var sentinel *NodeI
	var node *NodeI
	for sentinel = nil; sentinel != l.Next; sentinel = node {
		for node = l; node.Next != sentinel; node = node.Next {
			node.Data, node.Next.Data = node.Next.Data, node.Data
		}
	}
	return l
}
