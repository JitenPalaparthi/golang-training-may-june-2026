package main

type Square struct {
	S float32
}

func NewSqare(s float32) *Square {
	return &Square{s}
}

func (s *Square) Area() float64 {
	return float64(s.S * s.S)
}

func (s *Square) Perimeter() float64 {
	return 4 * float64(s.S)
}
