package main

import (
	"github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes/pkg/circle"
	"github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes/pkg/custom"
	"github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes/pkg/rect"
	"github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes/pkg/shape"
	"github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes/pkg/square"
)

func main() {
	r1 := rect.NewRect(98.67, 99.98)
	shape.Shape(r1)
	s1 := square.NewSqare(23.56)
	shape.Shape(s1)
	c1 := circle.NewCircle(23.56)
	shape.Shape(c1)
	//shape.printShape(c1) // cannot access the function, bcz it is restricted

	var t1 custom.T1

	t1.C = "Hello World"
	t1.Display()
	// t1.print

	t2 := custom.Newt2(12.23, 32.23, "Hello Wrold", true)
	t2.Display()

	shape.Greet()
}

// go get github.com/JitenPalaparthi/golang-personal-training-june-july-2026-shapes
