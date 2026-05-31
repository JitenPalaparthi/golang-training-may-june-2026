package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {

	var num1 uint8 = 250
	var num2 uint = uint(num1)
	println(num2) // no data is lost

	// all numbers can be casted amonng them
	// do they lose data ? If yes how do they

	var num3 uint = 0b11001100_11001100_11001100_11001100_11001100_11001100_11001100_11001100
	println(num3)
	var num4 uint8 = uint8(num3)

	var num5 uint = 14757395258967641292
	var num6 uint8 = uint8(num5)

	println(num4)
	println(num6)

	println(0b11001100)

	var float1 float32 = 12312.123 // IEEE-75X notation SIGN Exponent Fraction

	var num7 = uint(float1)
	println(num7)

	bits := math.Float32bits(float1)
	fmt.Printf("%032b\n", bits)

	ok1 := false // 1 byte

	//num8 = uint8(ok1) // cannot be type casted

	num8 := ToInt(ok1)
	println(num8)

	char1 := rune('A')

	fmt.Println(reflect.TypeOf(char1), char1)

	fmt.Println(string(char1)) // type caste to string
	char1++
	char1 = char1 + 'B'
	fmt.Println(reflect.TypeOf(char1), char1)

	var char2 uint64 = '世' // 19990

	var char3 uint64 = 19990
	str1 := string(char2)
	str2 := string(char3)
	println(str1, str2)
	num9 := 19990

	str3 := strconv.Itoa(num9) // type conversion

	fmt.Println("Type:", reflect.TypeOf(str3), "Value:", str3)

	a, b, c, d := calc(200, 100)
	println(a, b, c, d)

	a1, b1, _, _ := calc(20, 10) // _blank identifier

	println(a1, b1)

	_, _, _, d1 := calc(123, 12)
	println(d1)

	i, _ := 10, 20
	println(i) // not need to assign 20 to _ blank identifier, it is a valid expression

	str4, str5, str6 := "32131313321", "123121.12312", "14334d"

	val1, err := strconv.Atoi(str4)
	if err != nil { // that means there is an error
		println(err.Error())
	} else {
		println("successfully converted str4 to int", val1)
	}

	val2, _ := strconv.Atoi(str4) // can use blank identifier instead of err, when you ensure there is no error but not a good practice
	println(val2)

	va13, err := strconv.ParseFloat(str5, 64)

	if err != nil { // that means there is an error
		println(err.Error())
	} else {
		println("successfully converted str5 to float64", va13)
	}

	val4, err := strconv.ParseInt(str6, 10, 64)

	if err != nil { // that means there is an error
		println(err.Error())
	} else {
		println("successfully converted str4 to int64", val4)
	}

	str7 := "0a"
	val6, _ := strconv.ParseInt(str7, 10, 64)
	println(val6)

	str := fmt.Sprint(100, 200, true, " ", " hello Wrold ", any(231)) // everything can be converted to string using Sprint function
	println(str)

	// final

	{

		num1, num2, num3, num4, num5, num6, float1, float2, ok1, str1 := uint(12321), uint8(123), uint16(32312), -312312, int64(-213123), 4343, float32(6534.45), 5546454.456, true, "65454"
		sum := float64(0)
		sum = float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(num6) + float64(float1) + float2

		sum += float64(ToInt(ok1))

		v, _ := strconv.Atoi(str1) // okay with blank identifer, if it failes it returns 0 adding 0 has no issue

		sum += float64(v)
		fmt.Printf("sum=%.4f", sum)
	}

}

// type casting
// 1. implicit 2.explicit , Go has only explicit type casting
// type conversion
// type assertion

func ToInt(b bool) int { // Conversion function
	if b {
		return 1
	}
	return 0
}

func calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
