package piscine

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}
	newNode := &NodeL{Data: data, Next: nil}
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	l.Tail = newNode
}
