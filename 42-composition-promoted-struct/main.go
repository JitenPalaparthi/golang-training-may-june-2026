package main

import "fmt"

func main() {

	p1 := Person{Id: 100, Name: "Jiten", Email: "JitenP@Outlook,Com", Status: "active", Address: &Address{City: "Trivandrum", PinCode: "695011", Status: "current"}}

	fmt.Println(p1)

	fmt.Println("Pincode:", p1.Address.PinCode)
	fmt.Println("Address Status:", p1.Address.Status)

	e1 := Employee{Id: 100, Name: "Jiten", Email: "JitenP@Outlook,Com", Status: "active", Address: &Address{City: "Trivandrum", PinCode: "695011", Status: "current"}}
	fmt.Println("Pincode:", e1.PinCode)
	fmt.Println("Emoloyee Status:", e1.Status, "Address Status:", e1.Address.Status)

	//social := &SocialMedia{"linkedin.com/jpalaparthi", "linkedin.com/jpalaparthi", "X.com/jiten", "Insta.com/jiten"}

	e1.Linkedin = "linkedin.com/jpalaparthi"
	e1.Facebook = "linkedin.com/jpalaparthi"
	e1.Twitter = "X.com/jiten"
	//(*&e1).Instagram = "Insta.com/jiten"
	e1.Instagram = "Insta.com/jiten"

	//e1.SocialMedia = social
	fmt.Println(e1)

}

type Person struct {
	Id                  int
	Name, Email, Status string
	Address             *Address // composition
}

type Address struct {
	City, PinCode, Status string
}

type Employee struct {
	Id                  int
	Name, Email, Status string
	*Address            // promoted field, can relate it to Inheritance(not 100% but kind of )
	// *SocialMedia        // promoted field , dont use pointers
	SocialMedia
}

type SocialMedia struct {
	Linkedin  string
	Facebook  string
	Instagram string
	Twitter   string
}
