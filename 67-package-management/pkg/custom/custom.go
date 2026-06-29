package custom

import "fmt"

type T1 struct {
	a, b float32 // restricted
	C    string  // unrestricted
	Ok   bool    // unrestricted
}

func (t1 T1) print() { // restricted
	fmt.Println(t1)
}

func (t1 T1) Display() { // unrestricted
	fmt.Println(t1)
}

func Newt2(a, b float32, c string, ok bool) *t2 {
	return &t2{a, b, c, ok}
}

type t2 struct { // restricted
	a, b float32
	C    string
	Ok   bool
}

func (t2 t2) print() {
	fmt.Println(t2)
}

func (t2 t2) Display() {
	fmt.Println(t2)
}
