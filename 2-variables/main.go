package main

import (
	"fmt" // standard package fmt
	"reflect"
	"unsafe"
)

func main() {

	var num1 int // zero value for every variable

	var float1 float32

	var num2 uint8

	var ok1 bool // false

	var str1 string // empty string ""

	fmt.Println(num1, float1, num2, str1, ok1)

	num1 = 100

	ok1 = true

	str1 = "Hello World!"

	fmt.Println(num1, float1, num2, str1, ok1)

	var num3 int = 3213

	var float2 float64 = 3454355.34345

	num3++ // increment
	float2--

	fmt.Println(num3, float2) // defeloper friendly Println with log

	fmt.Printf("%d %f\n", num3, float2) // with proper decimals

	// type inference

	var num4 = 1321312 // int

	var float3 = 934324.34 // float64

	var str2 = "Hello Golang Folks!" //string

	var ok2 = true // bool

	fmt.Printf("num4:%d float3: %f str3 :%s ok2: %v\n", num4, float3, str2, ok2)

	var (
		num5   = 10 // int
		float4 = 1.35
	)

	fmt.Println("num5:", num5, "float4", float4)
	fmt.Println("Type of num5:", reflect.TypeOf(num5), "Size of num5:", unsafe.Sizeof(num5))
	//fmt.Println("Type of float4:", reflect.TypeOf(float4))
	fmt.Printf("Type of float4:%T\n", float4)

	//shorthand declaration

	a := 21312 // shorthand declaration
	b, c := 32.23, true

	fmt.Println("a", a, "b", b, "c", c)

	{
		var a = "Hello World"
		println(a)
	}

}

// variables and data types
// Numbers -> int, uint, int8,uint8,int16,uint16,int32,uint32,int64,uint64,float32,float64
// String -> string (collection of bytes)
// Bool -> bool
// Special Types -> any | interface{}

// zero value
