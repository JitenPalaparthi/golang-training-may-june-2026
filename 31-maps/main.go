package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func main() {

	slice := make([]int, 10)

	FillSlice(slice)

	fmt.Println(slice)
	// Find the repeated elements and also how many times a number is repeated in the sloice

	map1 := make(map[int]int)

	// for _, v := range slice {
	// 	// val, ok := map1[v]
	// 	// if ok {
	// 	// 	map1[v] = val + 1
	// 	// } else {
	// 	// 	map1[v] = 1
	// 	// }
	// 	map1[v] = map1[v] + 1
	// }

	FeedToMap(map1, slice)
	fmt.Println(map1)

	map1[999] = 19090
	fmt.Println(map1)
	delete(map1, 9999) // it deletes but is map is nil noops, if key does not exist no opes
	fmt.Println(map1)

	if err := Delete(map1, 999); err != nil {
		println(err.Error())
	} else {
		println("key successfully deleted from the map")
	}

	//delete(make(map[bool]int), true) // it does not have any elements, bcz map is created inside delete
}

func FillSlice(slice []int) {
	for i := range len(slice) {
		slice[i] = rand.IntN(10)
	}
}
func FeedToMap(map1 map[int]int, slice []int) {
	for _, v := range slice {
		map1[v] = map1[v] + 1
	}
}

func Delete(m map[int]int, key int) error {
	if m == nil {
		return errors.New("nil map")
	}
	if _, ok := m[key]; !ok {
		return fmt.Errorf("%v key does not exist", key)
	}

	delete(m, key)
	return nil
}
