package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// ptr ,len 10 cap 10

	fmt.Printf("slice1:%v Address of slice1:%p Ptr:%p Len:%d Cap:%d\n", slice1, &slice1, &slice1[0], len(slice1), cap(slice1))

	slice2 := slice1
	fmt.Printf("slice2:%v Address of slice2:%p Ptr:%p Len:%d Cap:%d\n", slice2, &slice2, &slice2[0], len(slice2), cap(slice2))

	slice3 := slice1[:5] // or slice1[0:5] 5 upto 5 but not 5
	fmt.Printf("slice3:%v Address of slice3:%p Ptr:%p Len:%d Cap:%d\n", slice3, &slice3, &slice3[0], len(slice3), cap(slice3))

	slice4 := slice1[3:8] // Len 5 Cap ?
	fmt.Printf("slice4:%v Address of slice4:%p Ptr:%p Len:%d Cap:%d\n", slice4, &slice4, &slice4[0], len(slice4), cap(slice4))
	// slice4[0] = 4444
	// slice4 = append(slice4, 9999)
	// fmt.Println(slice1)

	// slice4 = append(slice4, 1010, 1111, 1212)
	// slice4[1] = 5555
	// fmt.Println(slice1)

	slice5 := slice1[5:] // slice1[5:10]
	fmt.Printf("slice5:%v Address of slice5:%p Ptr:%p Len:%d Cap:%d\n", slice5, &slice5, &slice5[0], len(slice5), cap(slice5))

	// copy

	slice6 := make([]int, 5) // brand new slice with its own memory

	copy(slice6, slice1) // deep or element by elemnt copy

	fmt.Println(slice6)

	slice7 := []int{}

	Copy(slice7, slice1)
	fmt.Println(slice7)

	var slice8 []int

	Copy(slice8, slice1) //if dest slice is nil or len is 0 noops
	fmt.Println(slice8)

	slice9 := make([]int, 20) // right now zero values

	copy(slice9, slice1) // copies what ever is possible , the rest remains zero values

	fmt.Println(slice9)

	clear(slice9) // makes the whole slice as zero values, clear func does not make the slice as nil or lengh 0 just make all the elements in the slice as zero values
	fmt.Println(slice9)
	// for instance to delete[not a delete , to rearange] 3th index from slice1 and rearrange
	slice1 = append(slice1[:3], slice1[4:]...) // 0 to 3 , but not 3 index , 4th index till the end

	fmt.Println(">>>>>>> Len:", len(slice1), cap(slice1))
	// slice1 = append(slice1[0:3], 4444)
	// slice1 = append(slice1, slice2[4:]...)

	// [1:8] // 1 is the index , till 7th index but not 8
	//[1 2 3 5 6 7 8 9 10]
	fmt.Println(slice1)

	slice1 = append([]int{}, 50, 60, 70) // creates a new header and assigns to the existing slice1

	var slice10 []int // Ptr nil Len 0 Cap 0
	slice1 = append(slice10, 70, 80, 90)

	// slice11 := make([][]int, 10)
}

// [ 1 2 3 4 5 6 7 8 9 10]
// [4 5 6 7 8] Len 5 Cap 8

func Copy(dest, src []int) {
	min := min(len(dest), len(src))
	for i := 0; i < min; i++ {
		dest[i] = src[i]
	}
}

// Write a func to delete an element from the given slice and return that slice
