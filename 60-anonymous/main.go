package main

import (
	"fmt"
	"reflect"
)

func main() {

	sum := Calc(10, 20, func(i1, i2 int) int {
		return i1 + i2
	})

	println("sum:", sum)

	sub := Calc(20, 10, Sub)

	println("sub:", sub)

	var Fn func(int, int) int = func(i1, i2 int) int { return i1 * i2 }

	mul := Calc(20, 2, Fn)

	println("mul:", mul)

	var fn1 Func = func(i1, i2 int) int { return i1 / i2 }

	div := Calc(20, 10, fn1)
	println("div:", div)

	sq := fn1.Sq(10)
	println("sq:", sq)

	var calcMap map[string]any

	calcMap = make(map[string]any)

	calcMap["add"] = func(i1, i2 int) int {
		return i1 + i2
	}

	calcMap["sub"] = Sub

	calcMap["mul"] = Fn

	calcMap["div"] = fn1

	calcMap["greet"] = func() {
		println("Hello World")
	}

	calcMap["integer"] = 100
	calcMap["calc"] = Calc(100, 20, fn1)
	calcMap["done"] = true

	for k, v := range calcMap {
		fmt.Println(k, "-->", reflect.TypeOf(v))
	}

	for k, v := range calcMap {

		a, b := 20, 40
		switch vf := v.(type) {
		case func():
			println(k, ":")
			vf()
			println()

		case func(int, int) int:
			println(k, ":")
			r := vf(a, b)
			println("result:", r)
			println()

		case Func:
			println(k, ":")
			r := vf(b, a)
			println("result:", r)

			println("sq:", vf.Sq(10))

			println("Trying to type assert and then convert to a func(int,int)int which is generally not required")
			r2 := ((func(int, int) int)(vf))(10000, 20) // can do that but not required, directly access vf as a type Func and just execute it
			println("result:", r2)
			println()

		case int:
			println(k, ":", vf)
		default:
			println(k, ": ", "not implemented")
			println()
		}
	}
}

func Calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func Sub(a, b int) int {
	return a - b
}

type Func func(int, int) int

func (f Func) Sq(a int) int {
	return a * a
}

/*
greet --> func()
integer --> int
calc --> int
add --> func(int, int) int
sub --> func(int, int) int
mul --> func(int, int) int
div --> main.Func
*/
