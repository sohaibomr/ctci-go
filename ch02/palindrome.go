package ch02

//Implement a function to check if a linked list is a palindrome

//ABA, AOOA

// one way is to create a reversed linked list and compare both lists

func isPalindrome(list1 *LinkedList) bool {
	revList := newReversedList(list1)
	cur := list1.Head
	revCur := revList.Head
	for cur != nil {
		if cur.Val != revCur.Val {
			return false
		}
		cur = cur.Next
		revCur = revCur.Next
	}
	return true
}
