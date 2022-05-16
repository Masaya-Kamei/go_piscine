package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func getMaxNode(l *NodeI) (*NodeI, *NodeI) {
	var max_prev *NodeI = nil
	max := l
	var prev *NodeI = nil
	node := l
	for node != nil {
		if max.Data < node.Data {
			max_prev, max = prev, node
		}
		prev, node = node, node.Next
	}
	return max_prev, max
}

func ListSort(l *NodeI) *NodeI {
	var sorted *NodeI

	for l != nil {
		max_prev, max := getMaxNode(l)
		if max_prev == nil {
			l = max.Next
		} else {
			max_prev.Next = max.Next
		}
		max.Next = sorted
		sorted = max
	}
	return sorted
}
