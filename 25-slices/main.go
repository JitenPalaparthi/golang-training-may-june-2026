package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 5, 10)
	FillSlice(slice1)
	fmt.Println(slice1)

	slice2 := slice1 // create another slice from existing slice

	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	fmt.Printf("slice2:%v Address of slice2:%p Ptr:%p Len:%d Cap:%d\n", slice2, &slice2, &slice2[0], len(slice2), cap(slice2))

	slice2[0] = 1111

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	slice2 = append(slice2, 5555, 6666, 7777, 8888, 9999) // which header got appended

	// append 100% would change the header of the slice

	// append either changes len or changes all

	slice2[1] = 2222
	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	fmt.Printf("slice2:%v Address of slice2:%p Ptr:%p Len:%d Cap:%d\n", slice2, &slice2, &slice2[0], len(slice2), cap(slice2))

	slice2 = append(slice2, 101010) // Len:11 Cap 20
	slice2[2] = 3333

	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	fmt.Printf("slice2:%v Address of slice2:%p Ptr:%p Len:%d Cap:%d\n", slice2, &slice2, &slice2[0], len(slice2), cap(slice2))

	// What would happen to slice1?

	slice1 = slice2 // it is the copy of headers

	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))
	fmt.Printf("slice2:%v Address of slice2:%p Ptr:%p Len:%d Cap:%d\n", slice2, &slice2, &slice2[0], len(slice2), cap(slice2))

}

// Slice Header
// ptr , len , cal
func FillSlice(slice []int) {
	fmt.Printf("Address of slice:%p Ptr:%p Len:%d Cap:%d\n", &slice, &slice[0], len(slice), cap(slice))
	for i := range slice {
		slice[i] = rand.IntN(len(slice) * len(slice)) // just gave some range
	}
}

// 0x4691fdedc000
// 0x4691fdec8050
