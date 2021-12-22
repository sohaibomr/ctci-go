package ch02

func (linkedList *LinkedList) deleteNode(node *Node) {
	//	deleteNode deletes the given node from the linked list if it's not first or last node

	if linkedList.Head != node || node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}
