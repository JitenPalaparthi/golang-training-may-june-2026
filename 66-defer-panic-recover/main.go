package main

import (
	"fmt"
	"os"
)

func main() {

	fo := NewFileOps("")
	fmt.Fprintln(fo, "Hello World! How are you doing")

	fo1 := NewFileOps("abcd.txt")
	fmt.Fprintln(fo1, "Hello World! How are you doing")

	println("I am checking does it work, or reach here, since there is a panic inside the fileops")

}

type FileOps struct { // Concrete type
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

func (f *FileOps) Write(data []byte) (n int, err error) {
	//defer Re
	defer func() {
		if r := recover(); r != nil { // whe. would recover returns a value? when only there is a panic
			fmt.Println("I would be called when panic", r)
		}
		fmt.Println("There may be or may be not a panic, I would be called")
	}()

	if f == nil {
		panic("nil FileOps")
	}

	if f.FileName == "" {
		panic("invalid or empty filename")
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

func Recover() {
	if r := recover(); r != nil { // whe. would recover returns a value? when only there is a panic
		fmt.Println("I would be called when panic", r)
	}
}
