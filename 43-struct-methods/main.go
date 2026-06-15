package main

import "fmt"

func main() {

	r1 := Rect{L: 10.23, B: 12.34}

	//r2 := &Rect{L: 10.23, B: 12.34}

	println("call/pass by value")

	a1, p1 := Area(r1), Perimeter(r1)

	fmt.Printf("Area of Rect r1:%.2f\n", a1)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", p1)

	fmt.Printf("Area of Rect r1:%.2f\n", r1.A)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", r1.P)

	println("call/pass by ref")
	a1, p1 = AreaP(&r1), PerimeterP(&r1)

	fmt.Printf("Area of Rect r1:%.2f\n", a1)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", p1)

	fmt.Printf("Area of Rect r1:%.2f\n", r1.A)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", r1.P)

	println("call/pass by value as a method")

	a1, p1 = r1.Area(), r1.Perimeter()

	r1.L = 15.34
	r1.B = 34.34

	fmt.Printf("Area of Rect r1:%.2f\n", a1)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", p1)

	fmt.Printf("Area of Rect r1:%.2f\n", r1.A)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", r1.P)

	println("call/pass by ref by method")
	a1, p1 = (*&r1).AreaP(), r1.PerimeterP()

	fmt.Printf("Area of Rect r1:%.2f\n", a1)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", p1)

	fmt.Printf("Area of Rect r1:%.2f\n", r1.A)
	fmt.Printf("Perimeter of Rect r1: %.2f\n", r1.P)

	r2 := NewRect(123.34, 343.3)

	(*r2).L = 10.45 // dereferencing is Not required
	r2.B = 14.56

	println(r1.What())
	println(r2.What())

}

type Rect struct {
	L, B float32
	A, P float64
}

func NewRect(l, b float32) *Rect { // a consstructor kind of func
	return &Rect{L: l, B: b}
}

func Area(r Rect) float64 { // Area is a func, accespts an arg of Rect but that is call/pass by value
	r.A = float64(r.L) * float64(r.B)
	return r.A
}

func Perimeter(r Rect) float64 {
	r.P = 2 * (float64(r.L) + float64(r.B))
	return r.P
}

func AreaP(r *Rect) float64 { // Area is a func, accespts an arg of Rect but that is call/pass by value
	r.A = float64(r.L) * float64(r.B)
	return r.A
}

func PerimeterP(r *Rect) float64 {
	r.P = 2 * (float64(r.L) + float64(r.B))
	return r.P
}

// Methods are created using receivers
func (r Rect) Area() float64 { // Area is a func, accespts an arg of Rect but that is call/pass by value
	r.A = float64(r.L) * float64(r.B)
	return r.A
}

func (rectangle Rect) Perimeter() float64 { // using single letter as a receiver name is a convenction
	rectangle.P = 2 * (float64(rectangle.L) + float64(rectangle.B))
	return rectangle.P
}

func (*Rect) What() string {
	return "Rectangle"
}

func (r *Rect) AreaP() float64 { // Area is a func, accespts an arg of Rect but that is call/pass by value
	r.A = float64(r.L) * float64(r.B)
	return r.A
}

func (r *Rect) PerimeterP() float64 {
	r.P = 2 * (float64(r.L) + float64(r.B))
	return r.P
}
