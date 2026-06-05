package main

import "fmt"

func main() {

	slice1 := make([]int, 5, 10) // create and instatiaged a slice with len 5 and cap 10
	//slice1 = []int{100, 110, 12, 13, 14} // assigned values to the slice
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 10, 11, 12, 13, 14 // can assign like this as well
	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	slice1 = append(slice1, 15)
	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	slice1 = append(slice1, 16, 17, 18, 19, 20) // The len is more than the actual cap, it doubles the cap
	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	slice1 = append(slice1, 21, 22, 23, 24, 25) // The len is more than the actual cap, it doubles the cap
	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))

	// Ptr:0x367df54b8050 Len:5  Cap:10
	// Ptr:0x367df54b8050 Len:6  Cap:10
	// Ptr:0x367df54c2000 Len:11 Cap:20
	// Ptr:0x367df54c2000 Len:16 Cap:20

}
