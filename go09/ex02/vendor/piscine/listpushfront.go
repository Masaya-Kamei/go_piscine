package piscine

func ListPushFront(l *List, data interface{}) {
	if l == nil {
		return
	}
	newNode := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}
