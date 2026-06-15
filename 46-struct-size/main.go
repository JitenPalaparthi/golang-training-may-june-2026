package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var t1 T1
	var t2 T2
	var t3 T3

	fmt.Println("Size of t1:", unsafe.Sizeof(t1))
	fmt.Println("Size of t2:", unsafe.Sizeof(t2))
	fmt.Println("Size of t3:", unsafe.Sizeof(t3))
}

type T1 struct {
	i1 int   // 8
	b1 bool  // 1 8
	b2 bool  // 1  6 bytes are called padded bytes
	i2 int   // 8
	i3 int32 // 4 8
}

type T2 struct {
	i1     int   // 8
	i2     int   // 8
	i3     int32 // 8 - 4
	b1     bool  //
	b2     bool  //
	b3, b4 bool
}

type T3 struct {
	s1     string // 16 bytes as we assume but string itself is a structure --> inside that struct biggest size if 8 bytes
	i1     int    // 8
	i2     int    // 8
	i3     int32  // 8 - 4
	b1     bool   //
	b2     bool   //
	b3, b4 bool
}

// order of bytes
// padding
// optimization
