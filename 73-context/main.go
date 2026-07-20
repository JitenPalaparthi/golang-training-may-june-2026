package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	parent := context.Background() // creating a context

	ctx, cancel := context.WithDeadline(parent, time.Now().Add(time.Second*10))

	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				println("Goroutine is stoppedf")
				fmt.Println("Rerason:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("main received:", ctx.Err())

}

// type Context interface {
// 	Deadline() (deadline time.Time, ok bool)
// 	Done() <-chan struct{}
// 	Err() error
// 	Value(key any) any
// }
