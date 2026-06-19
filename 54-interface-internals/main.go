package main

import "fmt"

func main() {
	r1 := NewRect(98.67, 99.98)
	Shape(r1)
	s1 := NewSqare(23.56)
	Shape(s1)
	c1 := NewCircle(23.56)
	Shape(c1)

	var any1 any = c1
	any1 = s1
	any1 = r1
	any1 = 100
	any1 = "Hello World"
	fmt.Println(any1)

	var ie IEverything = c1
	ie = s1
	ie = r1
	ie = 100
	ie = "Hello World"
	fmt.Println(ie)

}

type IEverything interface{} // this is technically equal to any

// What technically is an interface ?
// In go irrespective of  any-> empty interface or any other interface like ishape are
// technically same

// any does any containes? there is nothing to implement, so every type is complaint to that interface
// interface{}

// every interface in go has dataptr and type pointer, whether it is any or ishape
// Shape is expecting any object that implements IShape interface
