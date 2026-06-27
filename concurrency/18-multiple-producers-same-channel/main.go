package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch, sig := make(chan string), make(chan struct{})

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go Produce(wg, ch, "Producer-1", 5)

	wg.Add(1)
	go Produce(wg, ch, "Producer-2", 10)

	wg.Add(1)
	go Produce(wg, ch, "Producer-3", 20)

	go func() {
		wg.Wait()
		close(ch)
	}()

	go Receiver(ch, sig)

	<-sig

}

func Produce(wg *sync.WaitGroup, ch chan string, name string, r uint) {
	defer wg.Done()
	for i := range r {
		time.Sleep(time.Millisecond * 100)
		ch <- fmt.Sprint(name, "--->", i)
	}
	// work is completed
	//close(ch)
}

func Receiver(ch chan string, sig chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			sig <- struct{}{}
			break
		} else {
			fmt.Println(v)
		}
	}
	// for v := range ch {
	// 	fmt.Println(v)
	// }
}
