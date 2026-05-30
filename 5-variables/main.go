package main

import (
	"fmt"
	"unsafe"
)

var (
	Global  int     // unassigned value stores in BSS segment
	Counter int = 9 // stores in Data Segment
)

const PI = 3.14

func main() {

	fmt.Println(Global, Counter, PI)

	fmt.Println(&Global, &Counter)

	a, b, c, d := 10, 20.3, true, "Hey" // tuple , there is no tuple in Go.
	println(a, b, c, d)

	/// generaly speaking string is a collection of chars , but in go string is collection of bytes

	str1 := "Hello World" // 11 bytes each and every char is an ascii chars
	str2 := "Hello, 世界"   // 9 bytes but not 13 bytes
	str3 := "Go has been available for a while now, backed by a big company and a growing community of users. It is advertised as a compiled, concurrent, imperative, structured programming language. It also has managed memory, so it looks safer and easier to use than C/C++. We have quite good experience with tools written in Go and decided to use it here. We have one open source project in Go, now we wanted to know how Go handles big traffic. We believed the whole project would take less than 100 lines of code and be fast enough to meet our requirements just because of Go."

	println("len of str1:", len(str1), "size of str1:", unsafe.Sizeof(str1))
	println("len of str2:", len(str2), "size of str2:", unsafe.Sizeof(str2))
	println("len of str3:", len(str3), "size of str3:", unsafe.Sizeof(str3))

	var str4 string // nothing has been given to string , still it is not nil, can check only "" (empty) or not

	if str4 == "" {
		println("it is empty string,probably the pointer inside the string header is nil, but we dont care --> ", len(str4))
	}

	str4 = "🚀" // emogis ,utf-8 char sequence

	println("Rocket char len:", len(str4))

	var ptr1 *bool   // 8 bytes
	var ptr2 *int64  // 8 bytes
	var ptr3 *uint16 // 8 bytes
	var ptr4 *string // 8 bytes

	if ptr1 == nil {
		println("ptr1 is nil")
	}

	println("size of ptr1:", unsafe.Sizeof(ptr1))
	println("size of ptr2:", unsafe.Sizeof(ptr2))
	println("size of ptr3:", unsafe.Sizeof(ptr3))
	println("size of ptr4:", unsafe.Sizeof(ptr4))

	// what  actually is string

	// string is a data structure -->
	/* stringHeader struct{
	   Ptr unsafe.Pointer // as of now assume it as a pointer // 8 bytes on 64 bit machine
	   Len int // 8 bytes
	}
	*/

	// when ever or where ever there is a pointer, that data type or the pointer can be nil (talking only abt the predefiend structures), but this rule does not apply to strings
	// there is no null in go , there is only nil

	// byte--> is not a special type in Go , it is just an alias for uint8 -> 0-255

	var b1 byte = 79
	var b2 uint8 = 79

	println(b1 + b2) // b1 and b2 are same

	// var a1 uint8 = 10
	// var b3 uint16 = 120

	// println(a1 + b3)

	var s1 singlechar = 123
	println(s1)

	// chars --> rune

	var char1 rune = 'A'
	var char2 int32 = '世'
	var char3 uint64 = '🚀'
	var char4 char = 'X'
	var char5 uint8 = 'a' //'世' //'a'

	println(char1, char2, char3, char4, char5)

	// rune is safe way to store 1-4 byte char
	// rune is for convenction, but rune is nothing more or less than int32 , it is just int32
	// can I store char more than 4 bytes, but you unnecessarily lose 4 bytes, becase the max size of any char is only 4 bytes

	bin1 := 0b1010 //0b a value in binary format
	hex1 := 0xffff // 0x a value in hexa format

	fmt.Printf("deciamail bin1:%d binary bin1: %b hexa bin1: %x\n", bin1, bin1, bin1)
	fmt.Printf("deciamail hex1:%d binary hex1: %b hexa hex1: %x\n", hex1, hex1, hex1)

	bin1 = bin1 + 20 // arthimetic opeartion on a value which is in the binary format, but the sytem does bin + deciaml no issues
	fmt.Printf("deciamail bin1:%d binary bin1: %b hexa bin1: %x\n", bin1, bin1, bin1)

	// bin1 = 0b1010
	// bin1 = bin1 | 0b1111
	// fmt.Println(bin1)

}

type singlechar = uint8 // singlechar is an alias for uint8, singlechar is not a new datatype

type char = int32 // creating char as an alias for int32

// 1. byte
// 2. rune.
// 3. any , interface{}
// 4. string
// 5. binary and hexa values

// 1001ac f50 B main.Global 0x105064 f50
// 10017c 468 D main.Counter 0x105034 468
