package main

import (
	"errors"
	"fmt"
)

// is it a concurrent quque? it is not , thread safe, how to make threadsafe queue
func main() {

	queue := NewQueue(10) //  0
	queue.Enqueue(20)     //  1
	queue.Enqueue(30)     //  2

	println("Current Index:", queue.CIndex)

	item := queue.Dequeue()
	fmt.Println("Current Index:", queue.CIndex, "Item:", item)

	item = queue.Dequeue()
	fmt.Println("Current Index:", queue.CIndex, "Item:", item)

	item = queue.Dequeue()
	fmt.Println("Current Index:", queue.CIndex, "Item:", item)

	item = queue.Dequeue()
	fmt.Println("Current Index:", queue.CIndex, "Item:", item)

	item = queue.Dequeue()
	fmt.Println("Current Index:", queue.CIndex, "Item:", item)

	queue.Enqueue(40) //  2

	item = queue.Dequeue()
	fmt.Println("Current Index:", queue.CIndex, "Item:", item)

}

type Queue struct {
	Data   []any
	CIndex int
}

func NewQueue(d any) *Queue {
	queue := append([]any{}, d)
	return &Queue{queue, 0}
}

func (q *Queue) Enqueue(item any) error {
	if q == nil || item == nil {
		return errors.New("nil queue or nil item")
	}
	q.Data = append(q.Data, item)
	//q.CIndex = q.CIndex + 1
	return nil
}

func (q *Queue) Dequeue() any {
	item := q.Data[q.CIndex]
	q.Data[q.CIndex] = nil
	if q.CIndex < len(q.Data)-1 {
		q.CIndex = q.CIndex + 1
	}
	return item
}
