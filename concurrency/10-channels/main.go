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
		wg.Done()
	}()

	go func() {
		for range 10 {
			println(<-ch)
			//<-ch
		}
		wg.Done()
	}()
	wg.Wait()
}

// send <-
// <- receive
