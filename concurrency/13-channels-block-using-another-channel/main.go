package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	//sig := make(chan bool) // why to spend 1 byte as bool when not doing anything with the value
	sig := make(chan struct{}) // sending a structure with out any memory,emoty structure
	go func() {
		for i := range 10 {
			time.Sleep(time.Millisecond * 100)
			ch <- i + 1
		}
		close(ch)

	}()

	go func() {
		for v := range ch { // range loop iterates until the channel is closed
			println(v)
		}
		//sig <- true
		sig <- struct{}{} // no data but can be a notification
	}()

	<-sig // future kind of , what are you doing with the value that is received?
}

// send <-
// <- receive
