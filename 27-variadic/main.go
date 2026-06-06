package main

import "fmt"

func main() {
	slice1 := append([]int{}, 10, 11, 12, 13, 14) // create a slice like this
	fmt.Println(slice1)

	var slice2 []int // a nil slice

	slice2 = append(slice2, 10) // can be appended so that it would get non nil ptr, len and cap

	slice3 := make([]int, 0, 5) // Len 2 Cap 5

	// slice3[0] = 1
	// slice3[1] = 2

	slice3 = append(slice3, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(slice3)

	//[0 0 3 4 5 6 7 8 9 10]

	// fmt.Println("a", "b", true, 10, 12, 30, "hey", 123.34)

	s := SumOf()
	println("Sum:", s)
	s = SumOf(10, 20)
	println("Sum:", s)
	s = SumOf(10, 20, 234, 45, 344, 34, 75, 56, 34, 35, 7, 45, 6576, 343, 765, 454, 345, 454)
	println("Sum:", s)
	s = SumOfMul(2, 20, 30, 40)
	println("Sum:", s)

	s = SumOf(slice1...) // pass a slice as a variadic arguments
	println("Sum:", s)

	arr1 := [5]int{10, 11, 12, 13, 14} // This is array

	arrslice := arr1[:] // This would be converted as slice

	// The ptr of array is nothginb but the 0th index pointeris given as ptr in slice header
	// the len of array is given as len and also as cap

	arrslice[0] = 143243
	fmt.Println(arr1)

	arrslice = append(arrslice, 10)
	arrslice[2] = 2222
	fmt.Println(arr1)

	s = SumOf(arr1[:]...)
	println("Sum:", s)

}

// variadic param
// can use range loop to access variadic args
// variadic param must be the last param
// variadic is only be used in funcs or methods cant be used as a variable or a field

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumOfMul(mul int, nums ...int) int {
	//func SumOfMul(nums ...int,mul int) int { // variadic param must be the last param
	sum := 0
	for _, v := range nums {
		sum += v * mul
	}
	return sum
}
