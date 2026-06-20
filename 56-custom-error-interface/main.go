package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello World!")
	// w io.Writer --> Interface

	fo := NewFileOps("data.txt")
	fmt.Fprintln(fo, "Hello World! How are you doing")

	var fo1 *FileOps // this is nil

	_, err := fmt.Fprintln(fo1, "Hello Go Folks, checking customized error")

	switch o := err.(type) {
	case *FileOpsError:
		println("Error Code:", o.Code)
		println("Error Message:", o.Msg)
	default:
		println(err.Error())
	}

	var fo2 = new(FileOps)

	_, err = fmt.Fprint(fo2, "Hello Golang minds, trying to implement error feature in Golang!")

	fileErr := new(FileOpsError)

	//var fileErr FileOpsError

	if errors.As(err, &fileErr) {
		println("Code:", fileErr.Code, "Message:", fileErr.Msg)
	} else {
		println("some error-->", err.Error())
	}
}

type FileOps struct { // Concrete type
	FileName string
}

func NewFileOps(fn string) *FileOps {
	return &FileOps{fn}
}

type FileOpsError struct {
	Code int
	Msg  string
}

func NewFileOpsError(code int, msg string) *FileOpsError {
	return &FileOpsError{code, msg}
}

func (f *FileOpsError) Error() string {
	return fmt.Sprint("Code:", f.Code, " ", "Message:", f.Msg)
}

// Error() string

func (f *FileOps) Write(data []byte) (n int, err error) {
	if f == nil {
		return 0, NewFileOpsError(101, "nil FileOps instance")
	}

	if f.FileName == "" {
		//return 0, fmt.Errorf("invalid filename:%v", f.FileName)
		return 0, NewFileOpsError(102, fmt.Sprintf("invalid filename:%v", f.FileName))
	}

	file, err := os.OpenFile(f.FileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return 0, NewFileOpsError(103, err.Error())
	}

	defer file.Close() // will discuss later about defer

	n, err = file.Write(data)
	if err != nil {
		return 0, NewFileOpsError(104, err.Error())
	}

	return n, nil
}

// error is an interface
// any type that implements that interface error can be an error
// for customized error handling , create types those implement error interface
// since error object is an interface, it can be type asserted to the original type.
