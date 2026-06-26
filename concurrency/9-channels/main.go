package main

import "time"

func main() {

	//var ch chan int     // declaration of channel , it is nil, bcz not instantiated
	ch := make(chan int) // buffered and unbuffered

	go func() {
		//time.Sleep(time.Second * 5)
		ch <- 100 // sending a value to the channel
		println("Value has been sent")
	}()

	time.Sleep(time.Second * 5)
	v := <-ch // blocked until it receives the value
	println(v)
	// -> = <-
}

// communication go uses channels
// channels are queue
// flow of data from one direction to the other.. it can never be bidirectional
// Why do we need channels?
// What kind of channels?

// Sender and also a receiver

// The size of the channel is very imp
// bcz when a sender sends the data to the channel, the sender can send only if the channel buffer is  free

// Block -->
// The sender is blocked until the receiver receives the value
// The receiver is blocked until the server sends the value
// There is a block against each other
