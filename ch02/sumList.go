package ch02

//Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.The digits are stored in reverse order,
//such that the 1 's digit is at the head of the list.
//Write a function that adds the two numbers and returns the sum as a linked list.
//EXAMPLE
//Input:(7-> 1 -> 6) + (5 -> 9 -> 2).That is,617 + 295. Output:2 -> 1 -> 9.Thatis,912.
//FOLLOW UP
//Suppose the digits are stored in forward order. Repeat the above problem. EXAMPLE
//lnput:(6 -> 1 -> 7) + (2 -> 9 -> 5).That is,617 + 295. Output:9 -> 1 -> 2.Thatis,912.

func sumList(list1 *LinkedListInt, list2 *LinkedListInt) *LinkedListInt {
	sumedList := newLinkedListInt()
	current1 := list1.Head
	current2 := list2.Head
	sum := 0
	reminder := 0
	for current1 != nil {
		sum += current1.Val + current2.Val
		nodeVal := sum
		if sum > 10 && current1.Next == nil { // last node

		} else if sum > 10 {
			reminder = sum % 10
			nodeVal = reminder
			sum = sum / 10

		}
		sumedList.addNode(&NodeInt{Val: nodeVal, Next: nil})
		current1 = current1.Next
		current2 = current2.Next
	}
	return sumedList

}
func valAndCarry(sum int) (int, int) {
	var val, carry int
	if sum < 10 {
		val = sum
		carry = 0
	} else if sum > 10 {
		val = sum % 10
		carry = sum / 10
	}
	return val, carry
}

//617, 295
func (linkedList *LinkedListInt) sumListReverse(node1 *NodeInt, node2 *NodeInt) int {
	if node1.Next == nil && node2.Next == nil {
		sum := node1.Val + node2.Val
		val, carry := valAndCarry(sum)
		linkedList.addNodeBefore(&NodeInt{
			Val:  val,
			Next: nil,
		})
		return carry
	}
	sum := node1.Val + node2.Val + linkedList.sumListReverse(node1.Next, node2.Next)
	val, carry := valAndCarry(sum)
	linkedList.addNodeBefore(&NodeInt{
		Val:  val,
		Next: nil,
	})
	return carry

}
