package ch02

//Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.

func (linkedList *LinkedList) kthToLast(k int) []string {
	current := linkedList.Head
	elems := []string{}
	i := 1
	for current.Next != nil {
		if i >= k {
			elems = append(elems, current.Val)
		}
		i++
		current = current.Next
	}
	elems = append(elems, current.Val)

	return elems
}
