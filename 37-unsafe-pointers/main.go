package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	//arr1 := [1000000]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//fmt.Println(arr1)
	fmt.Printf("Address of 0th element:%d\n", &arr1[0])
	fmt.Printf("Address of 1st element:%d\n", &arr1[1])
	fmt.Printf("Address of 3rd element:%d\n", &arr1[2])
	//fmt.Printf("Address of 1000000th element:%d\n", &arr1[1000000-1])

	var uptr uintptr = uintptr(unsafe.Pointer(&arr1[0]))

	// Exactly the pointer arthimetic usafe on arrays
	for range 9 { // from 1.22 onwards
		arrptr := (*int)(unsafe.Pointer(uptr))
		fmt.Println(*arrptr)
		uptr += unsafe.Sizeof(arr1[0])
	}

	println("call Iterate function")
	Iterate(unsafe.Pointer(&arr1[0]), 10, 8)

}

func Iterate(unsafePtr unsafe.Pointer, iter uint, size uint) {
	var uptr uintptr
	uptr = uintptr(unsafePtr)
	for range iter {
		v := (*int)(unsafe.Pointer(uptr))
		println(*v)
		uptr += uintptr(size)
	}
}
