package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var slice1 []any // Is it nil?

	slice2 := []any{} // Is it nil? Is it nil? not nil though len 0

	slice3 := make([]any, 0) // Is it nil? not nil though len 0

	var slice4 []any = []any{10, 20, true, "Hello World"} // not nil

	slice5 := make([]any, 0, 10) // Is it nil? not nil though len 0 and cap 10

	fmt.Println(slice1, slice2, slice3, slice4, slice5)

	// take any nil slice , like slice1 , how do you instantiate slice1

	slice1 = make([]any, 10) // Ptr: 0x318026d28640 Len: 10 Cap: 10
	fmt.Printf("Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", &slice1, &slice1[0], len(slice1), cap(slice1))
	FillSlice(slice1)

	// a := 10 // create in main stack frame
	// fmt.Printf("Outside addr of a %p\n", &a)
	// Incr(a) // pass a as argument , the value of a which is 10 is copied to a new value that is created for the functon Incr, call or pass by value
	// println(a)

}

func FillSlice(slice []any) {
	fmt.Printf("Address of slice:%p Ptr:%p Len:%d Cap:%d\n", &slice, &slice[0], len(slice), cap(slice))
	for i := range slice {
		slice[i] = rand.IntN(len(slice) * len(slice)) // just gave some range
	}
}

func Incr(a int) {
	fmt.Printf("Inside Address of a %p\n", &a)
	a++
}
