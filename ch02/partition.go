package ch02

func (linkedList *LinkedListInt) partition(x int) {
	var partionArr []int
	current := linkedList.Head
	for current != nil {
		if current.Val < x {
			partionArr = append(partionArr, current.Val)
		}
		current = current.Next
	}
	current = linkedList.Head
	for current != nil {
		if current.Val >= x {
			partionArr = append(partionArr, current.Val)
		}
		current = current.Next
	}
	current = linkedList.Head
	for _, val := range partionArr {
		current.Val = val
		current = current.Next
	}

}
