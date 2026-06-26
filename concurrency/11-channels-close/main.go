package main

import (
	"runtime"
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
		for {
			v, done := <-ch // same patten used fo any and also map
			if done {       // if true channel is not closed
				println(v)
			} else {
				wg.Done()
				runtime.Goexit()
				//break
			}
		}
	}()
	wg.Wait()
}

// send <-
// <- receive
