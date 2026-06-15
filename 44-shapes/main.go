package main

import "fmt"

func main() {

	r1 := NewRect(98.67, 99.98)
	a1, p1 := r1.Area(), r1.Perimeter()
	fmt.Printf("Area of Rect r1:%.2f\n", a1)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", p1)

	s1 := NewSqare(23.56)
	a2, p2 := s1.Area(), s1.Perimeter()
	fmt.Printf("Area of Squre s1:%.2f\n", a2)
	fmt.Printf("Perimeter of Square s1: %.2f\n", p2)

	c1 := NewCircle(23.56)
	a3, p3 := c1.Area(), c1.Perimeter()
	fmt.Printf("Area of Circle c1:%.2f\n", a3)
	fmt.Printf("Perimeter of Circle c1: %.2f\n", p3)
}
