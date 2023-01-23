package main

import "fmt"

type Queue struct {
	Items      []int64
	FrontIndex int64
	RearIndex  int64
}

func (q *Queue) New(size int64) {
	q.Items = make([]int64, size)
	q.FrontIndex = -1
	q.RearIndex = -1
}

func (q *Queue) IsEmpty() bool {
	return q.FrontIndex == -1
}

func (q *Queue) IsFull() bool {
	return int(q.RearIndex) == len(q.Items)-1
}

func (q *Queue) Enqueue(val int64) {
	if q.IsFull() {
		fmt.Printf("Queue is full!\n")
	} else {
		q.FrontIndex = 0
		q.RearIndex++

		q.Items[q.RearIndex] = val
		fmt.Printf("Enqueue %d!\n", val)
	}
}

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		fmt.Printf("Queue is empty!\n")
	} else {
		fmt.Printf("Dequeue %d!\n", q.Items[q.FrontIndex])
		q.FrontIndex++

		if q.FrontIndex == q.RearIndex {
			q.FrontIndex = -1
			q.RearIndex = -1
		}
	}
}

func (q *Queue) PrintQueue() {
	if q.IsEmpty() {
		fmt.Printf("Queue is empty!")
	} else {
		for i := q.FrontIndex; i <= q.RearIndex; i++ {
			fmt.Printf("%d ", q.Items[i])
		}
	}
	fmt.Printf("\n\n")
}

func main() {
	var q Queue

	q.New(3)
	q.PrintQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.PrintQueue()

	q.Dequeue()
	q.PrintQueue()

	q.Dequeue()
	q.PrintQueue()

	q.Dequeue()
	q.PrintQueue()

	q.PrintQueue()

	q.Enqueue(3)
	q.PrintQueue()

	q.Enqueue(2)
	q.PrintQueue()
}
