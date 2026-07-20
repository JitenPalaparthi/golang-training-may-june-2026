package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	parent := context.Background() // creating a context

	//ctx, cancel := context.WithDeadline(parent, time.Now().Add(time.Second*10))
	ctx, cancel := context.WithCancel(parent)

	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				println("Goroutine is stoppedf")
				fmt.Println("Rerason:", ctx.Err())
				runtime.Goexit()
				//return
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
