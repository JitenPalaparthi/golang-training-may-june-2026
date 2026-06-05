package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	var slice1 []int // declaration not instantiation

	if slice1 == nil {
		println("slice1 is nil")
		//slice1 = make([]int, 10) // make(Type,Len,Optional(Cap))
		slice1 = make([]int, 10, 20) // make(Type,Len,Optional(Cap))
		// Ptr: 0x510011 Len: 10 Cap: 20
	}

	fmt.Println(slice1)

	// becase we have not assgien any values, the slice has zero values

	FillSlice(slice1)

	fmt.Println(slice1)
	sum := Sumof(slice1)
	min, max, err := MinMax(slice1)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println("slice1--> Sum:", sum, "Max:", max, "Min:", min)
	}

	var slice2 []int

	sum = Sumof(slice2)
	min, max, err = MinMax(slice2)

	if err != nil {
		println(err.Error())
	} else {
		fmt.Println("slice2--> Sum:", sum, "Max:", max, "Min:", min)
	}

}

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(len(slice) * len(slice)) // just gave some range
	}
}

func Sumof(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func Max(slice []int) int {
	if len(slice) > 0 {
		max := slice[0]
		for _, v := range slice {
			if v > max {
				max = v
			}
		}
		return max
	}
	return 0
}

func MinMax(slice []int) (min int, max int, err error) {
	if slice == nil {
		return 0, 0, errors.New("nil slice")
	}
	if len(slice) > 0 {
		max, min = slice[0], slice[0]
		for _, v := range slice {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		return min, max, nil
	}
	return 0, 0, errors.New("slice has zero length")
}

// Arrays are fixed length
// Slices are not fixed length, extend
// Any slice of , e.g []int ,the type same for all int slices
// can write consistant funcs or methods, generally methods are written for slices
// an array can be converted as slice
// a slice can be nil , where as an array cannot be

// a slice can be nil --? Yes
// to instantiate a slice, use the built in function make

// make is used to create slice, map and chan
// the way make works is different for different types (slice,map and chan)

/* Slice Header
Ptr unsafe.Pointer
Len int
Cap int
*/
