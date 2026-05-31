package main

import (
	"fmt"
	"strconv"
)

func main() {

	num1, num2, num3, float1, float2 := 9444, uint64(3432423), uint16(43242), float32(676.45), float64(45345.345)
	var ok1 bool = true
	var str1 = "342342"
	var str2 = "1312.123"

	var any1 = any(4324.234)
	var any2 = any(42342)
	var any3 = any(ok1)

	var sum float64

	sum = float64(num1) + float64(num2) + float64(num3) + float64(float1) + float2
	sum += float64(ToInt(ok1))

	f1, _ := strconv.ParseFloat(str1, 64)
	//if err == nil {
	sum += f1
	//}
	i1, _ := strconv.Atoi(str2)
	sum += float64(i1)

	sum += any1.(float64)

	sum += float64(any2.(int))

	sum += float64(ToInt(any3.(bool)))

	fmt.Printf("sum:%.3f", sum)

}

func ToInt(b bool) int { // Conversion function
	if b {
		return 1
	}
	return 0
}
