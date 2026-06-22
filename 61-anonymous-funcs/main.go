package main

import "fmt"

func main() {

	fn1 := getSq(100)
	r := fn1()

	println("result:", r)

	funcSlice := make([]func(), 10)

	for i := range 10 {
		funcSlice[i] = func() {
			println("address:", &i, "sq:", i*i)
		}
	}

	fmt.Println(funcSlice)

	// am I executing the function --> No
	// since the funcs are in the slice , they would be executed at any point of time.

	for _, fn := range funcSlice {
		fn() // executing functions
		// 	println("r:", fn())
	}

	println("Same funcs stored in slice, but not thru for loop, using goto statement")
	i := 0
	funcslice1 := make([]func(), 10)

loop:
	funcslice1[i] = func() {
		println("address:", &i, "sq:", i*i)
	}
	i++
	if i < 10 {
		goto loop
	}

	for _, fn := range funcslice1 {
		fn() // executing functions
		// 	println("r:", fn())
	}
}

func getSq(num int) func() int {
	return func() int {
		return num * num
	}
}

// loops are optimized for the lazy init
// every iteration , the loop creats a new variable, which was not the case with the old versions of go
