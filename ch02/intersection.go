package ch02

func findIntersectingNode(list1 *LinkedListInt, list2 *LinkedListInt) *NodeInt {
	cur1 := list1.Head
	cur2 := list2.Head
	intersectingNode := &NodeInt{}
	for cur1 != nil && cur2 != nil {
		if cur1 == cur2 {
			intersectingNode = cur1
			break
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return intersectingNode
}
