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
