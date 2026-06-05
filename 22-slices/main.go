package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := []int{10, 20, 30, 40} // slice{Ptr:0x515500, Len:4 Cap: 4}

	// arr1 := [4]int{10, 20, 30, 40}
	// arr1 := [...]int{10, 20, 30, 40} // Automatically compiler evalues and put the lengh

	sum := Sumof(slice1)
	min, max, err := MinMax(slice1)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println("slice1--> Sum:", sum, "Max:", max, "Min:", min)
	}

	fmt.Println("slice1 Header", "Ptr:", &slice1[0], "Len:", len(slice1), "Cap:", cap(slice1))

	slice2 := []int{} // What is this slice? Is it nil , what is the len, cap

	slice3 := make([]int, 0) // not nil but len is zero and cap zero

	if slice2 == nil {
		println("yes, slice2 is nil")
	}
	if slice3 == nil {
		println("yes, slice3 is nil")
	}
	// This is not a nil slice
	// fmt.Println("slice1 Header", "Ptr:", &slice2[0], "Len:", len(slice2), "Cap:", cap(slice2))
	// Ptr:nil , Len:0 Cap :0

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

// To Instantiate the slice
// slice:=[]int{10,20} // not nil
// slice:=make([]int,2) // not nil
// slice:= []int{} // not nil

// Dummy Pointer or ZeroBase pointer --> The compiler creates a pointer which are used as Zerobase or dummy
//

// Class A {}
// Class B {}

// Class a = new(A)
// a.ToString()
// a.GetHashCode()
// class b = new(B)
