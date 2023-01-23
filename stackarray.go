package main

import "fmt"

type Stack struct {
	Items    []int64
	TopIndex int
}

func (s *Stack) New(size int64) {
	s.Items = make([]int64, size)
}

func (s *Stack) IsEmpty() bool {
	return s.TopIndex == 0
}

func (s *Stack) IsFull() bool {
	return s.TopIndex == len(s.Items)
}

func (s *Stack) Peek() {
	if s.IsEmpty() {
		fmt.Printf("Stack is empty!\n")
	} else {
		fmt.Printf("TopIndex item is: %d\n", s.Items[s.TopIndex-1])
	}
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Printf("Stack is empty!\n")
	} else {
		topVal := s.Items[s.TopIndex-1]
		s.Items[s.TopIndex-1] = -1
		s.TopIndex--
		fmt.Printf("Pop %d\n", topVal)
	}
}

func (s *Stack) Push(val int64) {
	if s.IsFull() {
		fmt.Printf("Stackoverflow\n\n")
	} else {
		s.Items[s.TopIndex] = val
		s.TopIndex++
		fmt.Printf("Push %d\n", val)
	}
}

func (s *Stack) Print() {
	if s.TopIndex == 0 {
		fmt.Println("Stack is empty!")
	} else {
		for i := s.TopIndex - 1; i > -1; i-- {
			fmt.Printf("%d\n", s.Items[i])
		}
	}
}

func main() {
	var s Stack
	var option int8
	var val int64
	var arrSize int64

	fmt.Printf("Input array size: ")
	fmt.Scanf("%d\n", &arrSize)
	s.New(arrSize)

	fmt.Printf("\n==============================\n")

	for {
		fmt.Printf("Choose option:\n")
		fmt.Printf("1. Peek\n")
		fmt.Printf("2. Pop\n")
		fmt.Printf("3. Push\n")
		fmt.Printf("4. Print\n")
		fmt.Printf("5. Exit!\n")
		fmt.Printf("Input option: ")
		fmt.Scanf("%d\n", &option)

		if option == 1 {
			s.Peek()
		} else if option == 2 {
			s.Pop()
		} else if option == 3 {
			fmt.Printf("Input value to push: ")
			fmt.Scanf("%d\n", &val)
			s.Push(val)
		} else if option == 4 {
			s.Print()
		} else if option == 5 {
			fmt.Printf("Exit!")
			break
		} else {
			fmt.Printf("Option %d is not available!", option)
		}

		fmt.Printf("\n==============================\n\n")
	}
}
