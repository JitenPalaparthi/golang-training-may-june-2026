package main

import "time"

func main() {

	//timeout := time.After(time.Second * 5)
	ch := Generate(time.Millisecond * 100)
	timeout := TimeOut(time.Second * 2)

	done := false // need to check to use or not to use
out:

	for {
		select {
		case v, ok := <-ch:
			if ok {
				println(v)
			} else {
				break out
			}
		case _, ok := <-timeout:
			if ok {
				if !done {
					done = true
					close(ch)
				}
			}
		}
	}
}

func Generate(du time.Duration) chan int {
	ch := make(chan int)
	go func() {
		i := 1
		for {
			time.Sleep(du)
			ch <- i * i
			i++
		}
	}()
	return ch
}

func TimeOut(duration time.Duration) <-chan struct{} {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(duration)
		timeout <- struct{}{}
		close(timeout)
	}()
	return timeout
}

// two channels
// 1- kept on sending data
// 2- say timeout
