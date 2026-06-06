package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice1)

	sum := AddElemAndSum(slice1, 6, 7, 8, 9, 10)

	fmt.Println("Sum:", sum)

	fmt.Println(slice1)

	sum, slice1 = AddElemAndSumR(slice1, 6, 7, 8, 9, 10)
	fmt.Println(slice1)

}

// when ever using append inside a func or a method , keep extra focus on the slice header
func AddElemAndSum(slice []int, elems ...int) (sum int) {
	// for _, e := range elems {
	// 	slice = append(slice, e)
	// }
	slice = append(slice, elems...)

	for _, v := range slice {
		sum += v
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
