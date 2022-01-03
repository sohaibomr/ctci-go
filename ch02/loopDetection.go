package ch02

func (linkedList LinkedList) detectLoop() bool {
	var slowNode *Node
	var fastNode *Node
	slowNode = linkedList.Head
	fastNode = linkedList.Head
	for fastNode.Next != nil {
		fastNode = fastNode.Next.Next
		if slowNode == fastNode {
			return true
		}
		slowNode = slowNode.Next
	}
	return false
}
