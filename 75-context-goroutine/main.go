package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"runtime"
)

func main() {

	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	ch := make(chan int)

	go Generate(ctx, ch, 7)
	go JustPrint(ctx)

	go func() {
		<-ch
		cancel()
	}()

	<-ctx.Done()

	fmt.Println("main received:", ctx.Err())

}

func Generate(ctx context.Context, ch chan int, stopDiv int) {
	for {
		select {
		case <-ctx.Done():
			println("Goroutine is stoppedf")
			fmt.Println("Rerason:", ctx.Err())
			runtime.Goexit()
		default:
			v := rand.IntN(10000)
			if v%stopDiv == 0 {
				ch <- v
			}
			if v%2 == 0 {
				println("Even Value:", v)
			} else {
				println("Odd Value:", v)
			}
		}
	}
}

func JustPrint(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			println("JustPrint goroutine is stoppedf")
			fmt.Println("Rerason:", ctx.Err())
			runtime.Goexit()
			//return
		default:
			fmt.Println("JustPrint goroutine is running")
			//time.Sleep(1 * time.Second)
		}
	}
}
