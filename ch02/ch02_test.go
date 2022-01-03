package ch02

import (
	"fmt"
	"testing"
)

func TestRemoveDups(t *testing.T) {

	linkedList := newLinkedList()
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "b", Next: nil})
	linkedList.addNode(&Node{Val: "b", Next: nil})
	linkedList.printList()
	fmt.Println("After removing dups")
	linkedList.removeDups()
	linkedList.printList()

	currNode := linkedList.Head
	for _, val := range []string{"a", "b"} {
		if currNode.Val != val {
			t.Fatal("Remove Dups failed! duplicate found")
		}
		currNode = currNode.Next
	}
}

func TestKThTOLast(t *testing.T) {
	linkedList := newLinkedList()
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "b", Next: nil})
	linkedList.addNode(&Node{Val: "b", Next: nil})
	expected := []string{"a", "b", "b"}
	actual := linkedList.kthToLast(4)
	fmt.Println(actual)
	fmt.Println(expected)
}

func TestDeleteMiddleNode(t *testing.T) {
	dNode := &Node{Val: "d", Next: nil}
	linkedList := newLinkedList()
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "b", Next: nil})
	linkedList.addNode(&Node{Val: "c", Next: nil})
	linkedList.addNode(dNode)
	linkedList.addNode(&Node{Val: "e", Next: nil})
	linkedList.addNode(&Node{Val: "f", Next: nil})
	linkedList.deleteNode(dNode)

	curNode := linkedList.Head
	for curNode.Next != nil {
		if curNode.Val == "d" {
			t.Fatalf("Deleted node should not be in the list")
		}
		curNode = curNode.Next
	}
}

func TestPartion(t *testing.T) {
	x := 5
	linkedList := newLinkedListInt()
	linkedList.addNode(&NodeInt{Val: 3, Next: nil})
	linkedList.addNode(&NodeInt{Val: 5, Next: nil})
	linkedList.addNode(&NodeInt{Val: 8, Next: nil})
	linkedList.addNode(&NodeInt{Val: 5, Next: nil})
	linkedList.addNode(&NodeInt{Val: 10, Next: nil})
	linkedList.addNode(&NodeInt{Val: 2, Next: nil})
	linkedList.addNode(&NodeInt{Val: 1, Next: nil})
	linkedList.partition(x)
	expected := []int{3, 2, 1, 5, 8, 5, 10}
	curr := linkedList.Head
	for _, val := range expected {
		if val != curr.Val {
			t.Fatalf("Linked list not partioned correctly")
		}
		curr = curr.Next
	}
}

func TestSumList(t *testing.T) {
	linkedList1 := newLinkedListInt()
	linkedList1.addNode(&NodeInt{Val: 7, Next: nil})
	linkedList1.addNode(&NodeInt{Val: 1, Next: nil})
	linkedList1.addNode(&NodeInt{Val: 6, Next: nil})

	linkedList2 := newLinkedListInt()
	linkedList2.addNode(&NodeInt{Val: 5, Next: nil})
	linkedList2.addNode(&NodeInt{Val: 9, Next: nil})
	linkedList2.addNode(&NodeInt{Val: 2, Next: nil})
	summedList := sumList(linkedList1, linkedList2)
	expected := []int{2, 1, 9}
	curr := summedList.Head
	for _, val := range expected {
		if val != curr.Val {
			t.Fatalf("Linked list not summed correctly")
		}
		curr = curr.Next
	}
}

func TestSumListReverse(t *testing.T) {
	revList := newLinkedListInt()
	linkedList1 := newLinkedListInt()
	linkedList1.addNode(&NodeInt{Val: 6, Next: nil})
	linkedList1.addNode(&NodeInt{Val: 1, Next: nil})
	linkedList1.addNode(&NodeInt{Val: 7, Next: nil})

	linkedList2 := newLinkedListInt()
	linkedList2.addNode(&NodeInt{Val: 2, Next: nil})
	linkedList2.addNode(&NodeInt{Val: 9, Next: nil})
	linkedList2.addNode(&NodeInt{Val: 5, Next: nil})

	revList.sumListReverse(linkedList1.Head, linkedList2.Head)
	expected := []int{9, 1, 2}
	curr := revList.Head
	for _, val := range expected {
		if val != curr.Val {
			t.Fatalf("Linked list not reversed sum correctly")
		}
		curr = curr.Next
	}
}

func TestPalindrome(t *testing.T) {

	linkedList := newLinkedList()
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "b", Next: nil})
	linkedList.addNode(&Node{Val: "c", Next: nil})
	linkedList.addNode(&Node{Val: "d", Next: nil})
	result := isPalindrome(linkedList)
	if result == true {
		t.Fatal("Palindrome failed! list is not palindrome")
	}
}

func TestPalindrome2(t *testing.T) {

	linkedList := newLinkedList()
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "o", Next: nil})
	linkedList.addNode(&Node{Val: "o", Next: nil})
	linkedList.addNode(&Node{Val: "a", Next: nil})
	result := isPalindrome(linkedList)
	if result == false {
		t.Fatal("Palindrome failed! list is a palindrome")
	}
}

func TestIntersectingNode(t *testing.T) {
	intersectingNode := &NodeInt{Val: 1, Next: nil}
	linkedList1 := newLinkedListInt()
	linkedList1.addNode(&NodeInt{Val: 6, Next: nil})
	linkedList1.addNode(intersectingNode)
	linkedList1.addNode(&NodeInt{Val: 7, Next: nil})

	linkedList2 := newLinkedListInt()
	linkedList2.addNode(&NodeInt{Val: 2, Next: nil})
	linkedList2.addNode(intersectingNode)
	linkedList2.addNode(&NodeInt{Val: 5, Next: nil})

	node := findIntersectingNode(linkedList1, linkedList2)
	fmt.Println(node)
	fmt.Println(intersectingNode)
}

func TestLoopDetection(t *testing.T) {

	linkedList := newLinkedList()
	loopNode := &Node{Val: "o", Next: nil}
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(&Node{Val: "o", Next: nil})
	linkedList.addNode(loopNode)
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNodeBefore(loopNode)
	result := linkedList.detectLoop()
	if result == false {
		t.Fatal("Loop detection failed")
	}
}

func TestLoopDetection2(t *testing.T) {

	linkedList := newLinkedList()
	loopNode := &Node{Val: "o", Next: nil}
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(loopNode)
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNodeBefore(loopNode)
	result := linkedList.detectLoop()
	if result == false {
		t.Fatal("Loop detection failed")
	}
}

func TestLoopDetection3(t *testing.T) {

	linkedList := newLinkedList()
	loopNode := &Node{Val: "o", Next: nil}
	linkedList.addNode(&Node{Val: "a", Next: nil})
	linkedList.addNode(loopNode)
	linkedList.addNodeBefore(loopNode)
	result := linkedList.detectLoop()
	if result == false {
		t.Fatal("Loop detection failed")
	}
}

func TestLoopDetection4(t *testing.T) {
	// loop detection with two same nodes

	linkedList := newLinkedList()
	loopNode := &Node{Val: "o", Next: nil}
	linkedList.addNode(loopNode)
	linkedList.addNodeBefore(loopNode)
	result := linkedList.detectLoop()
	if result == false {
		t.Fatal("Loop detection failed")
	}
}
