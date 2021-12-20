package ch02

//R�mov� Dups! Write code to remove duplicates from an unsorted linked list.
//FOLLOW UP
//How would you solve this problem if a temporary buffer is not allowed?

func (linkedList *LinkedList) removeDups() {
	var dups = make(map[string]bool)

	currentNode := linkedList.Head
	prevNode := linkedList.Head
	for prevNode.Next != nil {
		if _, ok := dups[currentNode.Val]; !ok {
			dups[currentNode.Val] = true
		} else {
			prevNode.Next = currentNode.Next
			currentNode = nil

		}
		if currentNode == nil {
			currentNode = prevNode.Next
		} else {
			prevNode = currentNode
			currentNode = currentNode.Next
		}
	}

}

// a -> a -> a -> a -> b -> b
