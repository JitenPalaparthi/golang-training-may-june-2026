package main

func main() {
	defer println("Hello World")
	// for i := range 10 {
	// 	defer println(rand.IntN(i + 1))
	// }
	defer fib(10)
	num := 0
	println(100 / num) // there is a panic, panics the call stack
	//fib(10)
}

func fib(r int) {
	println("fib numbers")
	a, b := 0, 1
	for range r {
		println(a)
		a, b = b, a+b
	}
}

// panic and defer works together
