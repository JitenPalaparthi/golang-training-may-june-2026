package main

import "time"

func main() {

	ch, sig := make(chan int), make(chan struct{})

	go Fib(ch, 10)

	go Rec("Worker-1", ch, sig)
	go Rec("Wroker-2", ch, sig)

	<-sig
	<-sig
}

func Fib(ch chan<- int, r int) { // Sender
	a, b := 0, 1
	for range r {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func Rec(worker string, ch <-chan int, sig chan<- struct{}) { // Receiver
	for c := range ch {
		time.Sleep(time.Millisecond * 100)
		println(worker, ">>>", c)
	}
	sig <- struct{}{}
}

// Future --> Sig is waiting for the value , until the value receivees sig blocks main
