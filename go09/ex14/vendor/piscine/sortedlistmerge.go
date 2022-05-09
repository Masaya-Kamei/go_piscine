package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func listBackClosure() func(*NodeI) *NodeI {
	var head *NodeI
	var tail *NodeI
	return func(n *NodeI) *NodeI {
		if n == nil {
			return head
		}
		if tail == nil {
			tail = n
			head = n
		} else {
			tail.Next = n
			tail = n
		}
		return head
	}
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var inserted *NodeI
	listBack := listBackClosure()

	for n1 != nil || n2 != nil {
		if n2 == nil || (n1 != nil && n1.Data <= n2.Data) {
			inserted = n1
			n1 = n1.Next
		} else {
			inserted = n2
			n2 = n2.Next
		}
		inserted.Next = nil
		listBack(inserted)
	}
	return listBack(nil)
}
