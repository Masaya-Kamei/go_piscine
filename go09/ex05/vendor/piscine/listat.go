package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	i := 0
	for node := l; node != nil; node = node.Next {
		if pos == i {
			return node
		}
		i++
	}
	return nil
}
