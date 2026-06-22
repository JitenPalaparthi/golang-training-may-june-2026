package main

import (
	"fmt"
	"os"
)

func main() {

	fo := NewFileOps("")
	fmt.Fprintln(fo, "Hello World! How are you doing")

	println("I am checking does it work, or reach here, since there is a panic inside the fileops")

}

type FileOps struct { // Concrete type
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

func (f *FileOps) Write(data []byte) (n int, err error) {
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
