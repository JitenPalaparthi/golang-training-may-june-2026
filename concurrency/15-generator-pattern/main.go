package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := Fib("sender-1", 10)
	ch2 := Fib("sender-2", 20)

	sig1 := Rec("Worker-1", ch1)
	sig2 := Rec("Worker-2", ch2)

	<-sig1
	<-sig2

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

func Rec(worker string, ch <-chan string) <-chan struct{} { // Receiver
	sig := make(chan struct{})
	go func() {
		for c := range ch {
			time.Sleep(time.Millisecond * 100)
			println(worker, ">>>", c)
		}
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}

// Future --> Sig is waiting for the value , until the value receivees sig blocks main
