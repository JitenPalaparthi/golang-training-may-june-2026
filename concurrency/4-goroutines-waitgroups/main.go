package main

import (
	"math/rand/v2"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup) // State=0

	wg.Add(1) // incr
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)))
			if i%2 == 0 {
				continue
			}
			println("Odd Goroutine:", i)
		}
		wg.Done() // dec the counter
	}()

	wg.Add(1)
	go func(r int) {
		a, b := 0, 1
		for range r {
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)))
			println("Fib Goroutine:", a)
			a, b = b, a+b
		}
		wg.Done()
	}(10)

	wg.Add(1)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(50)))
			if i%2 == 0 {
				println("Even Goroutine:", i)
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go iterateW(wg, 10)

	wg.Add(1)
	go func() {
		iterate(10)
		wg.Done()
	}()

	// wait would not consume cpu resources to wait, it uses a wakeup mechanism using sema internally by the runtime
	wg.Wait() // runtime_Semaacquire(&wg.sema), runtime_Semarelease(&wg.sema)
	// wait is a task

	// for {
	// 	if wg.State == 0 {
	// 		break
	// 	}
	// }

}

func iterate(n int) {
	for i := range n {
		time.Sleep(time.Microsecond * time.Duration(rand.IntN(50)))
		println("sequence in a func", i)
	}
}

func iterateW(wg *sync.WaitGroup, n int) {
	for i := range n {
		time.Sleep(time.Microsecond * time.Duration(rand.IntN(50)))
		println("sequence in a func with WG", i)
	}
	wg.Done()
}

// Similar to join handle --> WaitGroup
// WaitGroup maintains a state of a counter
