package main

func main() {
	r1 := NewRect(98.67, 99.98)
	Shape(r1)
	s1 := NewSqare(23.56)
	Shape(s1)
	c1 := NewCircle(23.56)
	Shape(c1)
}

// Shape is expecting any object that implements IShape interface
