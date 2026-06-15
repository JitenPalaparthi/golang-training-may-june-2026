package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// str := "Hello World"

	// bytes := []byte(str)

	// fmt.Println(string(bytes))

	fileOps := NewFileOps("data.txt")

	if err := fileOps.Write("Hello World, How are you doing!"); err != nil {
		println(err.Error())
	} else {
		println("data saved successfully")
	}

	if err := fileOps.Write("\nLearning a programming is a joy, if you feel it"); err != nil {
		println(err.Error())
	} else {
		println("data saved successfully")
	}

}

type FileOps struct {
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

// Write new methods , To read the file line by line and store in the []string , return it
func (fo *FileOps) Write(data string) error {

	if fo == nil {
		return errors.New("nil FileOps")
	}

	if fo.FileName == "" {
		return fmt.Errorf("invalid file name given:%v", fo.FileName)
	}

	f, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer f.Close() // will discuss later about defer

	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

func (fo *FileOps) ReadLines() ([]string, error) {

	if fo == nil {
		return nil, errors.New("nil FileOps")
	}

	if fo.FileName == "" {
		return nil, fmt.Errorf("invalid file name given:%v", fo.FileName)
	}

	f, err := os.OpenFile(fo.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	defer f.Close() // will discuss later about defer

	// for {
	// 	byte := make([]byte, 1)
	// 	_, err := f.Read(byte)
	// 	if err != nil {
	// 		break
	// 	}
	// }

	// TODO

	return nil, nil
}

// File/Directir -- 0
// R W E --> 4 2 1
// R W  --> 4 2 0 --> 6
// R --> 4 0 0
// O --> 4 0 0
// 0644
