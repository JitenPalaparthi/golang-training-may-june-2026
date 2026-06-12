package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int = 100
	var num2 int = 23423 //&num1
	//num1++

	//fmt.Printf("Address of num1 %b\n", &num1)
	fmt.Printf("Address of num1 %d\n", &num1)
	//fmt.Printf("Address of num1 %v\n", &num1)

	//fmt.Printf("Address of num2 %b\n", &num2)
	fmt.Printf("Address of num2 %d\n", &num2)
	//fmt.Printf("Address of num2 %v\n", &num2)
	//fmt.Printf("Address of num2 %o\n", &num2)

	// var ptr1 = &num1 // ptr1 is a pointer

	// fmt.Printf("Address that ptr1 holds:%d\n", ptr1)
	// println(ptr1)

	//ptr1 -= 84087011336232 // 0
	//00000000000000

	var uptr1 uintptr = uintptr(unsafe.Pointer(&num1)) // 1 & 4
	uptr1 += 8                                         // thinking that num2 is allocarted 8 bytes lesser than num1 so uptr1 may be pointed to num2

	v := (*int)(unsafe.Pointer(uptr1)) // 3 & 2
	fmt.Println(*v)

	fmt.Printf("Address of num1:%d\n", &num1)
	fmt.Printf("Address of num2:%d\n", &num2)
	fmt.Printf("Address of v:%d\n", v)

}

//84087011336232
//84087011336224

// only pointers can have the addressesm,even address is a number
// uintptr is specially designed to deal with unsafe.Pointer

// 1.A pointer value of any type can be converted to a unsafe.Pointer.
// 2.A unsafe.Pointer can be converted to a pointer value of any type.
// 3.A uintptr can be converted to a unsafe.Pointer.
// 4.A unsafe.Pointer can be converted to a uintptr.
