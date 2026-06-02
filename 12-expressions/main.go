package main

import "fmt"

func main() {

	a, b := 10, 20 // this assignment but also an expression

	_ = getSq(a) // This a func returns a value and so an expression, but stores the retur nvalue no where

	c := a + b + 10 + (a*a + b*b) + (a+b)/2 + (a + b%2) - 100 + 250/2*2

	// a + b + 10 + 500 + 15 + 10 - 100 + 250/2*2
	// a + b + 10 + 500 + 15 + 10 - 100 +125*2
	// a + b + 10 + 500 + 15 + 10 - 100 + 250
	// 715

	println(c)

	d := a+b+10+(a*a+b*b)+(a+b)/2+(a+b%2)-100+250/2*2 >= 700
	println(d)

	e := a+b+10+(a*a+b*b)+(a+b)/2+(a+b%2)-100+250/2*2 >= 720 && false || true

	// 715 >=700 && false || true
	// true && false || true
	// false || true
	// true

	println(e)

	b1 := uint8(0b10110110)

	// 7 6 5 4 3 2 1 0 position
	// 1 0 1 1 0 1 1 0 bits

	// 0 0 0 0 0 0 0 1
	// 0 0 1 0 0 0 0 0

	println(b1)

	mask := uint8(1 << 5)
	fmt.Printf("%08b, %08b\n", 1, mask)

	fmt.Printf("%08b\n", b1&mask)
	isOn := IsOn(b1 & mask)

	fmt.Printf("on:%v\n", isOn)

	// 7 6 5 4 3 2 1 0 position
	// 1 0 1 1 0 1 1 0 bits
	// 0 0 1 0 0 0 0 0
	// 0 0 1 0 0 0 0 0

}

func getSq(i int) int {
	return i * i
}

func IsOn(n uint8) bool {
	println(n)
	if n >= 1 {
		return true
	}
	return false
}

// expression -> what operators we give , arthmetic, comparision, logical and bitwise, assignment
// statement --> does not return just executes
