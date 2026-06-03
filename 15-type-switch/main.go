package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	var any1 any = 100

	any1 = 345.45

	fmt.Println("Value:", any1, "Type:", reflect.TypeOf(any1))

	//any1 = true

	fmt.Println(IsNumber(any1))
	// any3 = any1 + any2

	a, b := 10, 20

	if c, err := SumOf(a, b); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", c)
	}

	if c, err := SumOf(float32(123.32), float32(545.45)); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", c)
	}

	if c, err := SumOf(float32(123.32), float64(545.45)); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", c)
	}

	if c, err := SumOf(true, 123.23); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", c)
	}

	if c, err := SumOf("Hello", "World"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", c)
	}

}

func SumOf(a, b any) (any, error) {
	var sum = any(0.0)
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return sum, errors.New("input types are different")
	}

	if !IsNumber(a) {
		return sum, errors.New("inputarguments are not a number")
	}

	switch v := a.(type) { // v is a value directly returned
	case uint:
		sum = v + b.(uint)
	case uint8:
		sum = uint16(a.(uint8) + b.(uint8))
	case uint32:
		sum = uint64(a.(uint32) + b.(uint32))
	case uint64:
		sum = a.(uint64) + b.(uint64)
	case int:
		sum = a.(int) + b.(int)
	case int8:
		sum = int16(a.(int8) + b.(int8))
	case int16:
		sum = int32(a.(int16) + b.(int16))
	case int32:
		sum = int64(a.(int32) + b.(int32))
	case int64:
		sum = a.(int64) + b.(int64)
	case float32:
		sum = float64(a.(float32) + b.(float32))
	case float64:
		sum = a.(float64) + b.(float64)
	}

	return sum, nil
}

// func IsNumber(a any) bool {
// 	if _, ok := a.(int); !ok {
// 		if _, ok := a.(uint); !ok {
// 			if _, ok := a.(int8); !ok {
// 				if _, ok := a.(int16); !ok {
// 					if _, ok := a.(int32); !ok {
// 						if _, ok := a.(int64); !ok {
// 							if _, ok := a.(uint8); !ok {
// 								if _, ok := a.(uint16); !ok {
// 									if _, ok := a.(uint32); !ok {
// 										if _, ok := a.(uint64); !ok {
// 											if _, ok := a.(float32); !ok {
// 												if _, ok := a.(float64); !ok {
// 													return false
// 												}
// 											}
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return true
// }

func IsNumber(a interface{}) bool {
	switch a.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}
