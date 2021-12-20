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
