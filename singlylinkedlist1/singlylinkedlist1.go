// Implement Singly Linked List without Tail

package singlylinkedlist1

import "fmt"

type SinglyNode1 struct {
	Value int64
	Next  *SinglyNode1
}

type SinglyLinkedList1 struct {
	Head *SinglyNode1
}

func (sll *SinglyLinkedList1) InsertFirst(val int64) {
	var newNode SinglyNode1 = SinglyNode1{val, nil}

	if sll.Head == nil {
		sll.Head = &newNode
	} else {
		newNode.Next = sll.Head
		sll.Head = &newNode
	}
}

func (sll *SinglyLinkedList1) DeleteFirst() {
	if sll.Head == nil {
		fmt.Printf("Linked list is empty!\n")
	} else {
		deleteNode := sll.Head
		sll.Head = sll.Head.Next

		fmt.Printf("Value %d is deleted!\n", deleteNode.Value)
	}
}

func (sll *SinglyLinkedList1) InsertLast(val int64) {
	var newNode SinglyNode1 = SinglyNode1{val, nil}

	if sll.Head == nil {
		sll.Head = &newNode
	} else {
		currNode := sll.Head

		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = &newNode
	}
}

func (sll *SinglyLinkedList1) DeleteLast() {
	if sll.Head == nil {
		fmt.Printf("Linked list is empty!\n")
	} else {
		currNode := sll.Head
		var prevNode *SinglyNode1

		for currNode.Next != nil {
			prevNode = currNode
			currNode = currNode.Next
		}

		if currNode == sll.Head {
			sll.Head = nil
		} else {
			prevNode.Next = nil
		}

		fmt.Printf("Value %d is deleted!", currNode.Value)
	}
}

func (sll *SinglyLinkedList1) PrintLinkedList() {
	currNode := sll.Head

	if currNode == nil {
		fmt.Printf("Linked List is empty!\n")
	} else {
		for currNode != nil {
			fmt.Printf("%d ", currNode.Value)
			currNode = currNode.Next
		}
	}
}

func (sll *SinglyLinkedList1) Run() {
	var option int8
	var val int64

	fmt.Printf("\n==============================\n")

	for {
		fmt.Printf("Choose option:\n")
		fmt.Printf("1. Insert First\n")
		fmt.Printf("2. Insert Last\n")
		fmt.Printf("3. Delete First\n")
		fmt.Printf("4. Delete Last\n")
		fmt.Printf("5. Print\n")
		fmt.Printf("6. Exit!\n")
		fmt.Printf("Input option: ")
		fmt.Scanf("%d\n", &option)

		if option == 1 {
			fmt.Printf("Input value to insert first: ")
			fmt.Scanf("%d\n", &val)
			sll.InsertFirst(val)
		} else if option == 2 {
			fmt.Printf("Input value to insert last: ")
			fmt.Scanf("%d\n", &val)
			sll.InsertFirst(val)
		} else if option == 3 {
			sll.DeleteFirst()
		} else if option == 4 {
			sll.DeleteLast()
		} else if option == 5 {
			sll.PrintLinkedList()
		} else if option == 6 {
			fmt.Printf("Exit!")
			break
		} else {
			fmt.Printf("Option %d is not available!", option)
		}

		fmt.Printf("\n==============================\n\n")
	}
}
