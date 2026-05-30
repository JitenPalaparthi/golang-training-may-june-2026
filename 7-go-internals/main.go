package main

import (
	"fmt"          // standard library
	"math/rand/v2" // standard package
	"time"         // stadard package

	"github.com/google/uuid" // Third party package
)

func main() {

	fmt.Println("Hello World", time.Now())
	u := uuid.New()
	fmt.Printf("Generated UUID: %s\n", u.String())
	fmt.Println(rand.IntN(99999))

}

func greet() {
	println("Hello Golang folks!")
}

// go environment variables

// GOROOT --> It is a factory instalation of Go
// GOPATH --> It is a thirdpay library cashe, all thirdparty libraries or custom libraries come from here
// GOBIN  --> all tools, are installed in GOBIN, if GOBIN is empty the GOBIN location is $(GOPATH)/bin

// GOOS    -> Operating System name
// GOARCH  -> Chipset arch
