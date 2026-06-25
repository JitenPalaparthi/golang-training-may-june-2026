package main

import (
	"sync"
)

var (
	Counter = 0
	wg      = new(sync.WaitGroup)
)

func main() {

	for range 1000 {
		wg.Add(1)
		go Incr()
		wg.Add(1)
		go Decr()
	}

	wg.Wait()
	println(Counter)
}

// Do not communicate by sharing memory; instead, share memory by communicating.

func Incr() {
	Counter++
	wg.Done()
}

func Decr() {
	Counter--
	wg.Done()
}
