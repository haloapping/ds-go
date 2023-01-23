// Implement Singly Linked List without Tail

package main

import "fmt"

type SinglyNode1 struct {
	Value int64
	Next  *SinglyNode1
}

type SinglyLinkedList1 struct {
	Head *SinglyNode1
}

func (ll *SinglyLinkedList1) InsertFirst(val int64) {
	var newNode SinglyNode1 = SinglyNode1{val, nil}

	if ll.Head == nil {
		ll.Head = &newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = &newNode
	}
}

func (ll *SinglyLinkedList1) InsertLast(val int64) {
	var newNode SinglyNode1 = SinglyNode1{val, nil}

	if ll.Head == nil {
		ll.Head = &newNode
	} else {
		currNode := ll.Head

		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = &newNode
	}
}

func (ll *SinglyLinkedList1) PrintLinkedList() {
	currNode := ll.Head

	if currNode == nil {
		fmt.Printf("Linked List is empty!\n")
	} else {
		for currNode != nil {
			fmt.Printf("%d ", currNode.Value)
			currNode = currNode.Next
		}
	}
}

func main() {
	var ll SinglyLinkedList1

	ll.InsertLast(1)
	ll.InsertLast(2)

	ll.InsertFirst(3)
	ll.InsertFirst(4)

	ll.PrintLinkedList()
}
