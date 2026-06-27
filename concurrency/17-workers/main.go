package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := Fib("sender-1", 5, 50)
	sig := Workers(5, ch1)

	<-sig // it blocks until the sig is received

}

func Fib(name string, buf uint, r int) <-chan string { // Sender
	ch := make(chan string, buf) // buffered channel, since the buffered is given the producer is not blocked until the buffer is full
	go func() {
		a, b := 0, 1
		for range r {
			time.Sleep(time.Millisecond * 10)
			ch <- fmt.Sprint(name, ">>", a)
			a, b = b, a+b
		}
		close(ch)
	}()
	return ch
}

func Workers(workers uint, ch <-chan string) <-chan struct{} {
	sig := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(int(workers) + 1)

	go func() {
		wg.Wait()
		sig <- struct{}{}
		close(sig)
	}()

	go func() {
		for i := range workers {
			go func(i uint) {
				for c := range ch {
					time.Sleep(time.Millisecond * 50)
					println("Received by-->", i+1, ">>", c)
				}
				wg.Done()
			}(i)
		}
		wg.Done()
	}()

	return sig
}
