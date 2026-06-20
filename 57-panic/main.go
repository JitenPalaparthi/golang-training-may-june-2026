package main

func main() {

	num, div := 100, 1

	println(num / div)

	// div-- // div = 0 uncomment to panic

	println(num / div) // divide by zero is a very common mathamatical runtime error, which is panic here

	// arr := [5]int{10, 3, 4, 5, 76}

	// for i := 0; i <= len(arr); i++ { // paic index out of range
	// 	println(arr[i])
	// }

	var ptr *int //nil

	*ptr = 100 // dereferencing a nil pointer // nvalid memory address or nil pointer dereference

	println("Ptr:", *ptr) // this is also dereferencing a nil pointer

}

// panic
// panic panics the whole call stack, where ever there is a panic at that time, the application crashes
// panic is an error that cannot prodeed the exeecution of ther applicaiton
// an unhandled error

// runtime panic vs user defined panic
