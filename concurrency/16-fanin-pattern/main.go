package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := Fib("sender-1", 4)
	ch2 := Fib("sender-2", 5)
	ch3 := Fib("sender-3", 6)
	ch4 := Fib("sender-4", 7)
	ch5 := Fib("sender-5", 2)

	sig := FanInReceiver(ch1, ch2, ch3, ch4, ch5)

	<-sig // it blocks until the sig is received

}

func Fib(name string, r int) <-chan string { // Sender
	ch := make(chan string)
	go func() {
		a, b := 0, 1
		for range r {
			ch <- fmt.Sprint(name, ">>", a)
			a, b = b, a+b
		}
		close(ch)
	}()
	return ch
}

func FanInReceiver(chs ...<-chan string) <-chan struct{} {
	sig := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(len(chs) + 1) // add the number of channels those are passed as input args

	go func() {
		wg.Wait()
		sig <- struct{}{}
		close(sig)
	}()
	go func() {
		//time.Sleep(time.Second * 1)
		for i, ch := range chs {
			//wg.Add(1)
			go func() {
				for c := range ch {
					time.Sleep(time.Millisecond * 100)
					println("Worker-", i+1, ">>>", c)
				}
				//sig <- struct{}{}
				wg.Done()
			}()
		}
		wg.Done()
	}()
	return sig
}

// Future --> Sig is waiting for the value , until the value receivees sig blocks main
