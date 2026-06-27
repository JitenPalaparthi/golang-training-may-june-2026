package main

import (
	"fmt"
)

func main() {

	ch1 := Fib("sender-1", 4)
	ch2 := Fib("sender-2", 5)
	// ch3 := Fib("sender-3", 6)
	// ch4 := Fib("sender-4", 7)
	// ch5 := Fib("sender-5", 2)

	//con := make(chan string)

	//go func() {
	//count := 0
	done1, done2 := false, false
	for {
		//time.Sleep(time.Second * )
		// if count >= 5 {
		// 	break
		// }
		if done1 && done2 {
			break
		}
		select {
		case v, ok := <-ch1:
			if ok {
				println(v)
			} else {
				if !done1 {
					done1 = true
				}
			}
		case v, ok := <-ch2:
			if ok {
				println(v)
			} else {
				if !done2 {
					done2 = true
				}
			}
			// case v, ok := <-ch3:
			// 	if ok {
			// 		println(v)
			// 	} else {
			// 		count++
			// 	}
			// case v, ok := <-ch4:
			// 	if ok {
			// 		println(v)
			// 	} else {
			// 		count++
			// 	}
			// case v, ok := <-ch5:
			// 	if ok {
			// 		println(v)
			// 	} else {
			// 		count++
			// 	}
			// default:
			// 	println("Hit the default")
		}
	}
	//}()

}

func Fib(name string, r int) <-chan string { // Sender
	ch := make(chan string)
	go func() {
		a, b := 0, 1
		for range r {
			ch <- fmt.Sprint(name, ">>", a)
			a, b = b, a+b
		}
		close(ch)
	}()
	return ch
}

// Select only works on channels
// If multipel channels are trying to send or receive , select is a synchronization primitive
