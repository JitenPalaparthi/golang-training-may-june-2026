package main

import (
	"fmt"
)

func main() {

	a1 := add(10, 20) // add1(a,b int)int
	fmt.Println(a1)

	a1 = add(100, 200) // add1(a,b int)int
	fmt.Println(a1)

	a2 := add(float32(10.12), float32(12.34)) // add2(a,b float32)float32
	fmt.Println(a2)
	a3 := add(432423.234, 23423.234) // add3(a,b float64)float64
	fmt.Println(a3)

	// a3 := add(true, false)
	// fmt.Println(a3)

	// a4 := add(true, "Hello World")
	// fmt.Println(a4)
}

// func add(a, b int) int {
// 	return a + b
// }

// func add(a, b any) any {
// 	if reflect.TypeOf(a) == reflect.TypeOf(b) {
// 		switch a.(type) {
// 		case int:
// 			return a.(int) + b.(int)
// 		case uint:
// 			return a.(uint) + b.(uint)
// 		case float32:
// 			return a.(float32) + b.(float32)
// 			// default:
// 			// 	return nil
// 		}
// 	}
// 	return nil
// }

// func add[T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64](a, b T) T {
// 	return a + b
// }

func add[T INumber](a, b T) T {
	return a + b
}

type INumber interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

// monomorphization --> compiler generates teh code based on the call and rename the functions
// type eraser --> reverse concept , replaces types with any , any is nothing but an interface, an interface contains DataPtr and the TypePtr

// Go uses dynamic dispatch unlike monomorphization uses static dispatch

// polymorphism --> static dispatch
// int area(int s)
// float32 area(float32 a,float32 b)
// int a1 = area(100)
// float32 a2 = area(123.123,321.23)

// ==>
// int area1(int s)
// float32 area2(float32 a, float32 b)
// int a1 = area1(100)
// float32 a2 = area2(123.123,321.23)

// dynamic dispatch
// internally for every concrete type , go creates itable
// the instance and also the function pointer
// itable fof Rect instance ptr
// r1 area 0x1134 (function pointer for area)
// r1.area() --> what type is r1 --> it is rect
// where is the itable --> Go to that itable , look for instance r1,.
// look for function -- area --> pick the pointer and call it
