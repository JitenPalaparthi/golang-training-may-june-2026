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
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 100)
			println(i)
		}
	}()

	go greet()
	time.Sleep(time.Millisecond * 1100) // This is not a solution to block main by put it sleep
	// dont exit

}

func greet() {
	println("Hello World from greet")
}

// Main is also a goroutine
// No goroutine waits for othe goroutine to completes its execution
// if the main exists , the entire application exits
