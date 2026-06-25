package main

import (
	"math/rand/v2"
	"runtime"
	"time"
)

func main() {

	defer func() {
		// defer time.Sleep(time.Millisecond * 1000)
		defer runtime.Goexit()
		go func() {
			for i := 1; i <= 10; i++ {
				time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)))
				if i%2 == 0 {
					continue
				}
				println("Odd Goroutine:", i)
			}
		}()

	}()

	go func(r int) {
		a, b := 0, 1
		for range r {
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)))
			println("Fib Goroutine:", a)
			a, b = b, a+b
		}
	}(10)

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)))
			if i%2 == 0 {
				println("Even Goroutine:", i)
			}
		}
	}()

	//runtime.Goexit() //

	time.Sleep(time.Second * 2) // when I sleep it waits until the sleep is over so that the main would call the deffered calls

}

// Similar to join handle --> WaitGroup
//
