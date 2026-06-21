package main

import (
	"fmt"
	"unsafe"
)

func main() {

	func() { // func no input params, no return type
		// body of the func
		println("Hello World")
	}() // the executor

	c := func(a int, b int) int {
		return a + b
	}(10, 20) // c is the result of a+b

	println("sum:", c)

	f := func(r int) []int {
		println("addr:", &r)
		fib := make([]int, 0)
		a, b := 0, 1
		for i := 1; i <= r; i++ {
			fib = append(fib, a)
			a, b = b, a+b
		}
		// for i:= range r
		return fib
	}

	// f is a function not the result , reason is this function there is no executor so the whole funciton is stored in a variable
	// funcs are just variables in Go
	// funcs can be types in Go
	// the way we work with variables, same way can work with functions

	// some functions are called as inline functions by the compiler as and where compiler thinks that no need of a separate function..

	fib := f(10)
	fmt.Println(fib)

	var fn func(int, int) int = func(i1, i2 int) int {
		return i1 * i2
	}

	r := fn(40, 5)

	println("r:", r)

	fn = func(a int, b int) int {
		return a + b
	}

	r = fn(40, 5)
	println("r:", r)

	Greet()

	var fn2 func(int) []int

	if fn2 == nil {
		println("fn2 is nil func,size:", unsafe.Sizeof(fn2))
	}

	fn2 = f

	fmt.Printf("%p %p\n", fn2, f)

	fibr := fn2(5)

	fmt.Println(fibr)

	//fib2 := Fib(10)
	fmt.Printf("address of Fib:%p", Fib)
}

//go:noinline
func Greet() {
	println("Hey how are you doing!")
}

func Fib(r int) []int {
	println("addr:", &r)
	fib := make([]int, 0)
	a, b := 0, 1
	for i := 1; i <= r; i++ {
		fib = append(fib, a)
		a, b = b, a+b
	}
	// for i:= range r
	return fib
}
