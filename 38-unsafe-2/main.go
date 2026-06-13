package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//s1 := []int{} // Is it nil?  Len: 0 Cap: 0

	//var s1 []int // 0x0

	s1 := []int{}

	slicePtr := (*[3]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1))))

	ptr := (*int)(unsafe.Pointer(slicePtr[0]))
	//println(">>>>", &s1[0])
	fmt.Printf("ptr:%x\n", ptr)

	s2 := []int{}

	slicePtr = (*[3]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&s2))))

	ptr = (*int)(unsafe.Pointer(slicePtr[0]))
	//println(">>>>", &s1[0])
	fmt.Printf("ptr:%x\n", ptr)

	// 0x10064c ec0

	// 1001ac ec0

	var str string

	slicePtr1 := (*[2]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&str))))

	ptr = (*int)(unsafe.Pointer(slicePtr1[0]))
	//println(">>>>", &s1[0])
	fmt.Printf("ptr:%0Xx0\n", ptr)

	// if str == nil {

	// }

}
