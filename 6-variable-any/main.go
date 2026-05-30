package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var (
		num1   = 12312
		ok1    = true
		str1   = "Hello World"
		float1 = 1231.1233
	)

	// is there any type that can hold any kind of value of any data type?

	var any1 any // interface{}

	// what is the type and value of any

	fmt.Println("Value of any1", any1, "Type of any1:", reflect.TypeOf(any1))
	if any1 == nil { // can directly check any is nil or not
		println("Yes any1 is nil")
	}
	any1 = num1
	fmt.Println(any1, "size of any1", unsafe.Sizeof(any1))

	// int
	any1 = ok1
	// bool
	any1 = str1
	// string
	any1 = float1 // float64

	fmt.Println(any1, "size of any1", unsafe.Sizeof(any1))

}

/*
 AnyHeader
 ---------
 DataPointer unsafe.Pointer // 8
 TypePointer unsafe.Pointer // 8

*/
// Go is type safe programmin language
