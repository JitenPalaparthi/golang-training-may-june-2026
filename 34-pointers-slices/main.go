package main

import (
	"fmt"
	"unsafe"
)

func main() {

	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice1)

	sum := AddElemAndSum(slice1, 6, 7, 8, 9, 10)

	fmt.Println("Sum:", sum)

	fmt.Println(slice1) // new elements are not added to the original slice

	println("address of slice1:", &slice1)
	sum = AddElemAndSumPtr(&slice1, 6, 7, 8, 9, 10) // call or pass by reference
	fmt.Println("Sum:", sum)
	fmt.Println(slice1) // new elements are not added to the original slice

	// sum, slice1 = AddElemAndSumR(slice1, 6, 7, 8, 9, 10)
	// fmt.Println(slice1)

}

// when ever using append inside a func or a method , keep extra focus on the slice header
func AddElemAndSum(slice []int, elems ...int) (sum int) {
	// for _, e := range elems {
	// 	slice = append(slice, e)
	// }
	slice = append(slice, elems...) // changes the header

	for _, v := range slice {
		sum += v
	}
	return sum
}

// use pointer of slice and dereference
func AddElemAndSumPtr(ptr *[]int, elems ...int) (sum int) {
	println("Size of ptr:", unsafe.Sizeof(ptr), "address:", ptr)
	// for _, e := range elems {
	// 	slice = append(slice, e)
	// }
	if ptr != nil {
		*ptr = append(*ptr, elems...) // changes the header

		for _, v := range *ptr {
			sum += v
		}
	}
	return sum
}

func AddElemAndSumR(slice []int, elems ...int) (sum int, s []int) {
	// for _, e := range elems {
	// 	slice = append(slice, e)
	// }

	slice = append(slice, elems...)

	for _, v := range slice {
		sum += v
	}

	return sum, slice
}
