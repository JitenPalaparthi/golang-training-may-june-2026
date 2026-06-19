package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World!")

	fmt.Fprintln(os.Stdout, "Hello World!")
	// w io.Writer --> Interface

	fo := NewFileOps("data.txt")
	fmt.Fprintln(fo, "Hello World! How are you doing")

	// _, err := fmt.Fprintln(fo, "Hello World! How are you doing")
	// if err != nil {
	// 	println(err.Error())
	// }

}

type FileOps struct { // Concrete type
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

func (f *FileOps) Write(data []byte) (n int, err error) {
	if f == nil {
		return 0, errors.New("nil FileOps")
	}

	if f.FileName == "" {
		return 0, fmt.Errorf("invalid filename:%v", f.FileName)
	}

	file, err := os.OpenFile(f.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return 0, err
	}

	defer file.Close() // will discuss later about defer

	n, err = file.Write(data)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// interface is a contract, not an implementation, contract is all parties agree upon
// concrete types implements the interface
// to implement interface is implementing the definitions of the interface which are methods

// To design things, architecture
// It gives you a stable design, or a standard design
//
