package main

import (
	"fmt"
)

func main() {

	c1 := New(99, "red", "dark red")

	fmt.Println(c1)

	fmt.Println("Code:", c1.int)

	c2 := ColourCode{int: 3434, string: "Blue", FullColour: "Dark Blue"}
	fmt.Println(c2)

	// let t1 = (100,"Red","Dark Red")

	// w.r.t pointer they contain address
	// what address? Zerobase address
	e1 := new(Empty)
	n1 := new(Nothing)
	t1 := new(T)

	e2 := Empty{}

	fmt.Printf("e1 ptr:%p\n", e1)
	fmt.Printf("n1 ptr:%p\n", n1)
	fmt.Printf("t1 ptr:%p\n", t1)
	fmt.Printf("e2 ptr:%p\n", &e2)

}

type ColourCode struct {
	int        //anonymous fields
	string     // anonymous fields
	FullColour // since string is not allowd can create an alias for string and can use it here
}

type FullColour = string // not a new type but an alias

func New(c int, t string, fc FullColour) ColourCode {
	return ColourCode{c, t, fc}
}

type Empty struct{}   // an empty structure contains no fileds
type Nothing struct{} // an empty structure contains no fileds
type T struct{}       // an empty structure contains no fileds
