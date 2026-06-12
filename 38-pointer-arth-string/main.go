package main

import (
	"fmt"
	"unsafe"
)

func main() {

	str := "Hello World"

	fmt.Printf("Address of str:%d\n", &str)

	// barr := []byte(str)
	// fmt.Printf("Address of internal ptr:%d\n", &barr[0])

	// string is two elements of 8 bytes are strored togetner in a row like any other structure

	uptr := uintptr(unsafe.Pointer(&str))

	strHeader := (*[2]int)(unsafe.Pointer(uptr))

	(*strHeader)[1] = 1000 // it contains the len of the string,
	// I tampered the len of the string thru unsafe pointer

	//letPtr = 100
	println("length of str:", len(str))
	fmt.Println(str)

	slice1 := []int{10, 20, 30}

	fmt.Println("Len:", len(slice1), "Cap:", cap(slice1))

	uptr = uintptr(unsafe.Pointer(&slice1))

	sliceHeader := (*[3]int)(unsafe.Pointer(uptr))

	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	sliceHeader[0] = int(uintptr(unsafe.Pointer(&arr1)))
	sliceHeader[1] = 100
	sliceHeader[2] = 10101010

	fmt.Println("Len:", len(slice1), "Cap:", cap(slice1))
	fmt.Println(slice1)

	// var map1 map[string]any = map[string]any{"name": "jiten"}

	// fmt.Println(map1)

	// mapHeader := (*[10]int)(unsafe.Pointer(&map1))
	// //mapHeader[0] = 100000

	// fmt.Println("len:", len(map1), mapHeader)

}

/*
type hmap struct {
    count     int            // number of key/value pairs
    flags     uint8
    B         uint8          // log2 of number of buckets
    noverflow uint16         // approximate number of overflow buckets
    hash0     uint32         // hash seed

    buckets    unsafe.Pointer // pointer to array of buckets
    oldbuckets unsafe.Pointer // previous bucket array during grow
    nevacuate  uintptr        // progress counter for growing

    extra *mapextra           // optional fields for overflow handling
}

*/

// 71978440491120

// What is the internal header of string?

// string
// Ptr 8
// Len 8

// type Person struct {
// 	Id   int // Address if the first field
// 	Name string
// }
