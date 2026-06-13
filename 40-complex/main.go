package main

import (
	"fmt"
	"reflect"
)

func main() {

	c1 := 100.5 + 5.3i

	fmt.Println("Type if c1:", reflect.TypeOf(c1))

	c2 := complex(23, 2.4)

	fmt.Println(c1, c2)

	r, i := float32(32.3), float32(2.45)

	c3 := complex(r, i)

	fmt.Println("Type if c3:", reflect.TypeOf(c3))
	c4 := complex(float32(23), float32(2.4))

	c5 := c1 + c2

	c6 := c1 - complex128(c4)

	c7 := c1 * c2

	fmt.Println(c5, c6, c7)

	r1, i1 := real(c7), imag(c7)

	str := fmt.Sprintf("(%.2f+%.2f)", r1, i1)
	println(str)

	// This is a complex number

}

// complex128 and complex64
