// Implement Singly Linked List with Tail

package main

import "fmt"

type SinglyNode2 struct {
	Value int64
	Next  *SinglyNode2
}

type SinglyLinkedList2 struct {
	Head *SinglyNode2
	Tail *SinglyNode2
}

func (ll *SinglyLinkedList2) InsertFirst(val int64) {
	var newNode SinglyNode2 = SinglyNode2{val, nil}

	if ll.Head == nil {
		ll.Head = &newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = &newNode
	}
}

func (ll *SinglyLinkedList2) InsertLast(val int64) {
	var newNode SinglyNode2 = SinglyNode2{val, nil}

	if ll.Head == nil {
		ll.Head = &newNode
		ll.Tail = ll.Head
	} else {
		ll.Tail.Next = &newNode
		ll.Tail = &newNode
	}
}

func (ll *SinglyLinkedList2) PrintLinkedList() {
	currNode := ll.Head

	if currNode == nil {
		fmt.Printf("Linked List is empty!\n")
	} else {
		for currNode != ll.Tail.Next {
			fmt.Printf("%d ", currNode.Value)
			currNode = currNode.Next
		}
	}
}

func main() {
	var ll SinglyLinkedList2

	ll.InsertLast(1)
	ll.InsertFirst(2)
	ll.InsertLast(3)
	ll.InsertLast(10)
	ll.InsertFirst(-1)

	ll.PrintLinkedList()
}
