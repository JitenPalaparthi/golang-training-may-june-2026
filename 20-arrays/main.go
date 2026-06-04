package main

import (
	"fmt"
	"math/rand/v2"
	"unsafe"
)

func main() {

	arr1 := [5]int{} // zero values

	fmt.Println("Len:", len(arr1), "Size:", unsafe.Sizeof(arr1), "Data:", arr1)

	for i := range arr1 {
		arr1[i] = rand.IntN(100) // assigning rand numbers in arr
	}

	fmt.Println("Len:", len(arr1), "Size:", unsafe.Sizeof(arr1), "Data:", arr1, "Max:", Max(arr1))
	min, max := MinMax(arr1)
	fmt.Println("Min:", min, "Max:", max)

	println()
	arr2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2d)

	sum := 0
	for _, a1 := range arr2d {
		for _, v := range a1 {
			sum += v
		}
	}
	println("Sum:", sum)

	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Println(arr3d)

	// [[[0 0 0] [0 0 0]] [[0 0 0] [0 0 0]]]

	sum = 0

	for i := 0; i < len(arr3d); i++ {
		for j := 0; j < len(arr3d[i]); j++ {
			for k := 0; k < len(arr3d[i][j]); k++ {
				sum += arr3d[i][j][k]
			}
		}
	}

	println("sum of arr3d using normal nested loops", sum)
	sum = 0
	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v := range a2 {
				sum += v
			}
		}
	}
	println("sum of arr3d using range loops", sum)
}

func Max(arr [5]int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func MinMax(arr [5]int) (min int, max int) {
	max, min = arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}
