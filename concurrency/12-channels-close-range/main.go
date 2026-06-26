package main

import (
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		for i := range 10 {
			time.Sleep(time.Millisecond * 100)
			ch <- i + 1
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		for v := range ch { // range loop iterates until the channel is closed
			println(v)
		}
		wg.Done()
	}()
	wg.Wait()
}

// send <-
// <- receive
