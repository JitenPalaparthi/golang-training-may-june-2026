package main

import (
	"errors"
	"fmt"
)

func main() {

	var map1 MyMap

	map1 = make(MyMap)

	//map1 = make(map[string]any)

	map1["560086"] = "Blr-1"
	map1["560096"] = "Blr-2"
	map1["560034"] = "Blr-3"
	map1["695011"] = "Trv-1"
	map1["522001"] = "Guntur-1"
	map1["522002"] = "Guntur-2"

	println("map1\n------------------------")

	map1.Display("Pincode", "Area")

	if err := map1.Delete("522004"); err != nil {
		println(err.Error())
	} else {
		println("succesfully deleted from the map")
	}

	keys1, values1 := map1.GetKVPairs()

	fmt.Println("Keys:", keys1)
	fmt.Println("Values:", values1)

	println("map2\n----------------")
	var map2 map[string]any

	map2 = make(map[string]any)

	map2["560086"] = "Blr-1"
	map2["560096"] = "Blr-2"
	map2["560034"] = "Blr-3"
	map2["695011"] = "Trv-1"
	map2["522001"] = "Guntur-1"
	map2["522002"] = "Guntur-2"

	MyMap(map2).Display("Pincode", "Area")

	if err := MyMap(map2).Delete("522004"); err != nil {
		println(err.Error())
	} else {
		println("succesfully deleted from the map")
	}

	keys2, values2 := MyMap(map2).GetKVPairs()

	fmt.Println("Keys:", keys2)
	fmt.Println("Values:", values2)

	var map3 map[string]string

	map3 = make(map[string]string)

	map3["560086"] = "Bangalore-1"

	//MyMap(map3).Display() // This does not work bcz the underlining type is map[string]any not map[string]string

	slice1 := make(MySlice, 5)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 10, 20, 30, 40, 50

	slice1.Display()

	if max, err := slice1.Max(); err != nil {
		println(err.Error())
	} else {
		println("max:", max)
	}

	if min, err := slice1.Min(); err != nil {
		println(err.Error())
	} else {
		println("min:", min)
	}

	slice2 := make([]int, 5)
	slice2[0], slice2[1], slice2[2], slice2[3], slice2[4] = 310, 230, 306, 340, 506

	ms1 := MySlice(slice2)

	(&ms1).Display()

	//(&MySlice(MySlice(slice2))).Display()

	//(&MySlice(slice2)).Display()

}

type MyMap map[string]any

func (m MyMap) Display(keyName, valueName string) {
	for k, v := range m {
		fmt.Println(keyName, ":", k, valueName, ":", v)
	}
}

func (m MyMap) Delete(key string) error {
	if m == nil {
		return errors.New("nil map")
	}
	if _, ok := m[key]; !ok {
		return fmt.Errorf("%v key does not exist", key)
	}

	delete(m, key)
	return nil
}

func (m MyMap) GetKVPairs() (keys []string, values []any) {
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

type MySlice []int

func (s *MySlice) Display() {
	for _, v := range *s {
		fmt.Print(v, " ")
	}
}

func (s MySlice) Max() (int, error) {
	if s == nil {
		return 0, errors.New("nil map")
	}
	if len(s) == 0 {
		return 0, errors.New("no elements in the map,len is 0")
	}
	max := s[0]
	for _, v := range s {
		if max > v {
			max = v
		}
	}
	return max, nil
}

func (s MySlice) Min() (int, error) {
	if s == nil {
		return 0, errors.New("nil map")
	}
	if len(s) == 0 {
		return 0, errors.New("no elements in the map,len is 0")
	}
	min := s[0]
	for _, v := range s { // the underlining type is []int
		if min < v {
			min = v
		}
	}
	return min, nil
}
