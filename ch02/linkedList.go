package ch02

import (
	"fmt"
)

// Linked List

type Node struct {
	Val  string
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func newLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}
func newLinkedListWithElems(elems []string) *LinkedList {
	linkedList := newLinkedList()
	for _, elem := range elems {
		linkedList.addNode(&Node{Val: elem, Next: nil})
	}
	return linkedList
}
func (linkedList *LinkedList) addNode(node *Node) {

	if linkedList.Head == nil {
		//	first element
		linkedList.Head = node
	} else {
		currenNode := linkedList.Head
		for currenNode.Next != nil {
			currenNode = currenNode.Next
		}
		currenNode.Next = node
	}

}

func (linkedList *LinkedList) printList() {
	//TODO: write preety print for lists
	//var sb strings.Builder
	currentNode := linkedList.Head

	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}

}

func (linkedList *LinkedList) addNodeBefore(node *Node) {
	if linkedList.Head == nil {
		linkedList.Head = node
	} else {
		node.Next = linkedList.Head
		linkedList.Head = node
	}

}

func newReversedList(list *LinkedList) *LinkedList {
	revList := newLinkedList()
	cur := list.Head
	for cur != nil {
		revList.addNodeBefore(&Node{Val: cur.Val,
			Next: nil,
		})
		cur = cur.Next
	}
	return revList
}

// Linked List Int

type NodeInt struct {
	Val  int
	Next *NodeInt
}

type LinkedListInt struct {
	Head *NodeInt
}

func newLinkedListInt() *LinkedListInt {
	return &LinkedListInt{
		Head: nil,
	}
}
func (linkedList *LinkedListInt) addNode(node *NodeInt) {

	if linkedList.Head == nil {
		//	first element
		linkedList.Head = node
	} else {
		currenNode := linkedList.Head
		for currenNode.Next != nil {
			currenNode = currenNode.Next
		}
		currenNode.Next = node
	}

}

func (linkedList *LinkedListInt) printList() {
	//TODO: write preety print for lists
	//var sb strings.Builder
	currentNode := linkedList.Head

	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}

}

func (linkedList *LinkedListInt) addNodeBefore(node *NodeInt) {
	if linkedList.Head == nil {
		linkedList.Head = node
	} else {
		node.Next = linkedList.Head
		linkedList.Head = node
	}

}
func reversePrint(node *NodeInt) {
	if node != nil {
		reversePrint(node.Next)
		fmt.Println(node.Val)
	}

}
