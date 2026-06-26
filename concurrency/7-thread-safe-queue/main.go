package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {

	queue := NewQueue(10)
	// fmt.Println("CIndex:", queue.CIndex, "Data:", queue.Data)
	// queue.Enqueuq(20)
	// fmt.Println("CIndex:", queue.CIndex, "Data:", queue.Data)
	// queue.Enqueuq(30)
	// fmt.Println("CIndex:", queue.CIndex, "Data:", queue.Data)
	// queue.Enqueuq(40)
	// fmt.Println("CIndex:", queue.CIndex, "Data:", queue.Data)

	// item := queue.Dequeue()

	// fmt.Println("CIndex:", queue.CIndex, "Item:", item, "Data:", queue.Data)

	// item = queue.Dequeue()

	// fmt.Println("CIndex:", queue.CIndex, "Item:", item, "Data:", queue.Data)

	// item = queue.Dequeue()

	// fmt.Println("CIndex:", queue.CIndex, "Item:", item, "Data:", queue.Data)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		for i := range 10 {
			time.Sleep(time.Millisecond * time.Duration(100+rand.IntN(100)))
			fmt.Println("Enqueee>>") //"Data:", queue.Data)
			queue.Enqueuq((i + 1) * 10)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for range 20 {
			time.Sleep(time.Millisecond * time.Duration(100+rand.IntN(100)))
			fmt.Println("Dequee", "Item:", queue.Dequeue())
		}
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println("Dequee", "CIndex:", queue.CIndex, "Item:", queue.Dequeue(), "Data:", queue.Data)
	fmt.Println("Dequee", "CIndex:", queue.CIndex, "Item:", queue.Dequeue(), "Data:", queue.Data)
}

type Queue struct {
	Data   []any
	mu     *sync.Mutex
	CIndex int
}

func NewQueue(d any) *Queue {
	queue := append([]any{}, d)
	return &Queue{queue, new(sync.Mutex), 0}
}

func (q *Queue) Enqueuq(item any) error {
	if q == nil || item == nil {
		return errors.New("nil queue or item to insert")
	}
	q.mu.Lock()
	q.Data = append(q.Data, item)
	q.mu.Unlock()
	return nil
}

func (q *Queue) Dequeue() any {
	if len(q.Data) > q.CIndex {
		q.mu.Lock()
		item := q.Data[q.CIndex]
		q.Data[q.CIndex] = nil
		q.CIndex++
		q.mu.Unlock()
		return item
	}
	return nil
}

//|10| 20 |30 |40 |50 | 60|

// go uses dlv debugger
// c++ and rust used GDB and LLDB
