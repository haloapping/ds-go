package main

import "fmt"

type DoublyNode struct {
	Value int64
	Prev  *DoublyNode
	Next  *DoublyNode
}

type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
}

func (dll *DoublyLinkedList) InsertFirst(val int64) {
	newNode := DoublyNode{val, nil, nil}

	if dll.Head == nil {
		dll.Head = &newNode
		dll.Tail = dll.Head
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = &newNode
		dll.Head = &newNode
	}
}

func (dll *DoublyLinkedList) InsertLast(val int64) {
	newNode := DoublyNode{val, nil, nil}

	if dll.Head == nil {
		dll.Head = &newNode
		dll.Tail = dll.Head
	} else {
		dll.Tail.Next = &newNode
		newNode.Prev = dll.Tail
		dll.Tail = &newNode
	}
}

func (dll *DoublyLinkedList) PrintFromFirst() {
	currNode := dll.Head

	for currNode != dll.Tail.Next {
		fmt.Printf("%d ", currNode.Value)
		currNode = currNode.Next
	}
	fmt.Printf("\n")
}

func (dll *DoublyLinkedList) PrintFromLast() {
	currNode := dll.Tail

	for currNode != dll.Head.Prev {
		fmt.Printf("%d ", currNode.Value)
		currNode = currNode.Prev
	}
	fmt.Printf("\n")
}

func main() {
	var dll DoublyLinkedList

	dll.InsertFirst(1)
	dll.InsertFirst(2)
	dll.InsertFirst(3)
	dll.InsertFirst(4)
	dll.InsertLast(5)
	dll.InsertLast(6)
	dll.InsertFirst(10)
	dll.PrintFromFirst()
	dll.PrintFromLast()
}
