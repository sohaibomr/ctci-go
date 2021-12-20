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
