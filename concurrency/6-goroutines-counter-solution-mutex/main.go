package main

import (
	"sync"
)

var (
	Counter = 0 // By design global variables are not thread safe
	wg      = new(sync.WaitGroup)
	mu      = new(sync.Mutex)
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
	mu.Lock()
	Counter++
	mu.Unlock()
	wg.Done()
}

func Decr() {
	mu.Lock()
	Counter--
	mu.Unlock()
	wg.Done()
}
