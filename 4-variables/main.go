package main

var Counter = 1 // package level variable

const PI = 3.14 // RO

var Global int // unassigned , BSS

func main() {

	Counter++
	println(Counter)
	Incr()
	Incr()
	Incr()
	println(Counter)

	a, b, ok, str := 10, 20, true, "Hello World" // Where these variables are allocated

	println(a, b, ok, str)

	// wherther a variable is stored in stack or heap is decided by compiler, escape analysis

	// var arr1 [10]int
	// var arr2 [3123213213213]int
	// fmt.Println(len(arr1), "addr:", &arr1[0])
	// fmt.Println(len(arr2), "addr:", &arr2[0])

	t := b
	b = a
	a = t

	println(a, b)

	a, b = b, a
	println(a, b)

	a1, b1, c1 := 10, 20, 30

	a1, b1, c1 = b1, c1, a1

	println("a1:", a1, "b1:", b1, "c1:", c1)
}

/*
Code/Text Segment          --> The entire binay is loaded into this memory
Data Segment -> BSS, RO etc --> Global/package (also called as static) variables,consts are allocated here

Stack Memory --> 2MB(general standard)
Heap Memory  --> Unlimited amount of memory during runtime[Depends on availability and other OS factors]
*/
