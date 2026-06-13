package main

import "fmt"

func main() {

	var p1 Person

	p1.Id = 100
	p1.Name = "Jiten"
	p1.Email = "JitenP@Outllok.Com"

	fmt.Println(p1)

	p2 := Person{100, "Jiten", "JitenP@Outllok.Com", ""}
	fmt.Println(p2)

	p3 := Person{Id: 100, Email: "JitenP@Outllok.Com", Name: "Jiten"}
	p3.Status = "active"
	fmt.Println(p3)

	p4 := &Person{} // pointer

	(*p4).Id = 111 // no need to acccess with * operator
	p4.Name = "Jiten"
	p4.Email = "JitenP@Outlook.Com"
	p4.Status = "active"
	fmt.Println(p4)

	p5 := new(Person)

	(*p5).Id = 111 // no need to acccess with * operator
	p5.Name = "Jiten"
	p5.Email = "JitenP@Outlook.Com"
	p5.Status = "active"
	fmt.Println(p5)

	p6 := New(100, "Jiten", "JitenP@Outllok.Com")

	fmt.Println(p6)

	if p1 == p2 { // since the equal operator on struct, a struct can be a key to the map as well
		println("Yes p1 and p2 are equal")
	}

	mapPerson := make(map[Person]any) // ? because what can be the key , any type that can apply == operator

	fmt.Println(mapPerson)

}

type Person struct {
	Id                  int
	Name, Email, Status string
}

// Kind of a constructor
func New(id int, name, email string) *Person {
	return &Person{id, name, email, "active"}
}

// no classes only structs
// can do oops programming using struts-> not like OOPS programming language but in a Go's idiamatic way
// struct is a composite type or user defined type
// probably stored on stack memory (again depends on escape analaysis)

// No constrors , No polymorphism, No encapsulation (not in a traditional oops way), no Inheritance
// no public,private,protected,internal, friend etc.. kind of access specifiers in Go
// A method in go is a function that is attached to a type
