package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	if l == nil {
		return 0
	}
	size := 0
	for node := l.Head; node != nil; node = node.Next {
		size++
	}
	return size
}
