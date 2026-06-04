package main

import (
	"fmt"
	"unsafe"
)

var (
	arr5 = [...]int{34, 343, 45, 23, 76, 658, 4, 545, 45, 3, 76, 54} // The length is evaluated during compilation
)

func main() {

	var arr1 [3]int // there are zero values
	fmt.Println("len:", len(arr1), arr1)

	arr1[0] = 100               // mutating
	arr1[1], arr1[2] = 200, 300 // mutating
	//arr1[4] = 400
	fmt.Println("len:", len(arr1), arr1)

	var arr2 [5]int = [5]int{10, 20, 30, 40, 50} // creating and assigning values to the array
	fmt.Println("len:", len(arr2), arr2)

	arr3 := [6]int{10, 20, 30, 40, 60, 70} // shorthand declartion , create and assign values to the array

	fmt.Println("len:", len(arr3), arr3)

	arr4 := [...]int{34, 343, 45, 23, 76, 658, 4, 545, 45, 3, 76} // The length is evaluated during compilation
	fmt.Println("len:", len(arr4), arr4)

	fmt.Printf("Address of arr4:%p type of arr4:%T\n", &arr4, arr4)
	fmt.Printf("Address of arr5:%p type of arr5:%T\n", &arr5, arr5)

	// type of an array

	// var num1 uint8 = 143
	// var num2 uint16 = 143

	// if uint16(num1) == num2 {

	// }

	// This does not compile becase both are different types and cant do type casting like above

	// if arr2 == ([5]int)(arr3) {
	// 	println("yes both are equal")
	// } else {
	// 	println("No, both are different")
	// }

	var arr6 [5]any // what is the zero value of arr6
	// zero value of any any is nil

	fmt.Println(arr6)

	arr6[0] = 5433
	arr6[1] = "Hello World"
	arr6[2] = arr4
	arr6[3] = true
	arr6[4] = arr5

	fmt.Println(arr6, "Len:", len(arr6), "Size:", unsafe.Sizeof(arr6))

	var arr7 [5]any // even there are no data given

	fmt.Println(arr7, "Len:", len(arr7), "Size:", unsafe.Sizeof(arr7))

	if arr6 == arr7 {
		println("arr6 and arr7 both len and ofcourse size are same, values are also same")
	} else {
		println("arr6 and arr7 both len and ofcourse size are same,but values are not same")
	}

	arr8, arr9 := [3]int{10, 20, 30}, [3]int{20, 40, 60}

	//arr10 := arr8 + arr9

	arr10 := SumOfArrays(arr8, arr9)
	fmt.Println(arr10)

	sum := SumOf(arr2)
	println("sum:", sum)

	sum = SumOf6(arr3)

	// arr11, arr12 := [4]int{10, 20, 30, 40}, [4]int{20, 40, 60, 80}

	// arr13 := SumOfArrays(arr11, arr12) // types mismatch

	// var num1, num2 = float32(10.1), float32(20.2)
	// num3 := num1 + num2
	// println(num3)
	// arth(+,-,/%) --> uint,int,int8,int16,int32,int64,uint8,uint16,uint32,uint64,flaot32,float64
	// [1]int,[2]float64,[34324234234]float32
}

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOf6(arr [6]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// SIMD --> Single Instruction Multiple Data, just for understanding, here it does not do SIMD
func SumOfArrays(arr1, arr2 [3]int) [3]int {
	var arr3 [3]int
	for i := 0; i < len(arr1); i++ {
		arr3[i] = arr1[i] + arr2[i]
	}
	return arr3
}

// arrays sequence of elements of the same type
// arrays are fixed length
// array is immutable, w.r.t the size not with in the bounds
