package main

const PI = 3.14

type Circle struct {
	R float32
}

func NewCircle(s float32) *Circle {
	return &Circle{s}
}

func (s *Circle) Area() float64 {
	return PI * float64(s.R*s.R)
}

func (r *Circle) Perimeter() float64 {
	return 2 * PI * float64(r.R)
}
