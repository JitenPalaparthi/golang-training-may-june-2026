package main

import (
	"fmt"
	"hash/fnv"
	"math/rand/v2"
)

func main() {

	var m map[string]any // this is only declaration but not instantiation

	m = make(map[string]any, 10) // going to create buckets, how many to start with 2

	m["560086"] = "Banglaore-1"

	h := Hash64("560086" + fmt.Sprint(rand.IntN(100))) // some seed to look like real hash
	fmt.Println(h)
	fmt.Println(h % 16)

	m["560096"] = "Bangalore-2"
	h = Hash64("560096" + fmt.Sprint(rand.IntN(100))) // some seed to look like real hash
	fmt.Println(h)
	fmt.Println(h % 16)

	m["560034"] = "Bangalore-3"
	m["522001"] = "Guntur-1"
	m["522002"] = "Guntur-2"

	// fetch all value from the map
	// range loop gives , key and value

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	// to fetch a specific element

	v, ok := m["560086"]
	if ok {
		fmt.Println("Value:", v)
	} else {
		println("Key does not exist")
	}

	// elems := [4]int{10, 11, 12, 13}
	// if elems[0] == 12 {
	// 	println("0th element has 12")
	// } else if elems[1] == 12 {
	// 	println("1th element has 12")
	// } else if elems[2] == 12 {
	// 	println("2nd element has 12")
	// } else if elems[3] == 12 {
	// 	println("3nd element has 12")
	// } else {
	// 	println("element not found")
	// }

	// number eleentns/ number of buckets --> 10 /2 ->5
	// 13/2 --> 6.5
	// 14 --> 4 , two buckets is doubled to 4 buckets
	// rebalance all the elemnents in that map across 4 buckets, old bucket

	// What data type can be a key while creating a map -?
	// any data type that can use == operator can be key data type

	// m1 := make(map[[]int]string)     // it does not work
	// m2 := make(map[[2]uint64]string) // it does not work

	// s1 := []int{10, 20}
	// s2 := []int{10, 20}

	// if s1 == s2 { // This operation cannot be performed, hence a slice cannot be a key

	// }

	// s1 := [2]int{10, 20}
	// s2 := [2]int{10, 20}
	// if s1 == s2 { // no error so [2]int can be a key

	// }

}

func Hash64(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

// 239077177 6764849493
// Bucket no Index in bucket
// Go <1.24

// Go uses to create few buckets, each bucket contains 8 elements
// Each bucket contins parallel arrays
// Map -->< Bucket-0 Bucket-1
// Overflow bucket

// Loadfactor -> number of buckets / total number of elements -> 6.5, the buckets are doubled
// 238980200/%buvkets 7508956616%8 -> collissions
