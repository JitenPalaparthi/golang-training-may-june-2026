package main

import (
	"runtime"
	"time"
)

func main() { // main is G // main is runnign on P0

	procs := runtime.GOMAXPROCS(0) // On my suystem, it has 8 cores, so 8 Processes are created

	println("Number of Proces:", procs)

	go func() {
		println("Hello World-1")
	}()

	go func() {
		println("Hello World-2")
	}()

	go func() {
		println("Hello World-3")
		go func() {
			for i := 1; i <= 10; i++ {
				time.Sleep(time.Millisecond * 100)
				println("In another goroutine", i)
			}
		}()
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond * 100)
			println(i)
		}
	}()

	go func() {
		i := 1
		for {
			time.Sleep(time.Millisecond * 100)
			println(i)
			i++
			if i > 5 {
				runtime.Goexit()
			}
		}
	}()

	go greet()
	//time.Sleep(time.Millisecond * 1100) // This is not a solution to block main by put it sleep
	// dont exit
	runtime.Goexit()
	// Goexit 2 things 1. Is it running in any other goroutine, if yes, it waits for other goroutines to complete and exit.
	// 2. If it is called in main, it waits for all other goroutines to complete execution and crashes

	// So if call runtime.Goexit in main , you should know that your application would crash

	//time.Sleep(time.Second * 2)
}

func greet() {
	println("Hello World from greet")
}

// Main is also a goroutine
// No goroutine waits for othe goroutine to completes its execution
// if the main exists , the entire application exits
