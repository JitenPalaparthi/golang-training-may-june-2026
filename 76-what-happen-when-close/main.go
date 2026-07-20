package main

import (
	"sync"
	"time"
)

func main() {

	ch := make(chan int)

	wg := new(sync.WaitGroup)
	wg.Add(4)
	defer wg.Wait()

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- i
		}

		close(ch)
		wg.Done()
	}()

	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					println(v, ok)
				} else {
					wg.Done()
					return
				}
			}
		}

	}()

	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					println(v, ok)
				} else {
					wg.Done()
					return
				}
			}
		}

	}()

	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					println(v, ok)
				} else {
					wg.Done()
					return
				}
			}
		}

	}()

}
