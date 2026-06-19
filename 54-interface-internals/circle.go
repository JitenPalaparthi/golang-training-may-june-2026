package main

const PI = 3.14

type Circle float32

func NewCircle(s float32) Circle {
	return Circle(s)
}

func (s Circle) Area() float64 {
	return PI * float64(s*s)
}

func (r Circle) Perimeter() float64 {
	return 2 * PI * float64(r)
}

func (Circle) What() string { return "Circle" }
