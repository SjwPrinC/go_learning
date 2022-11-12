package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := errors.New("my error")

	fmt.Printf("err type is %T \n", err)
	// we can define our types by impl error interface in builtin file

	// error assert
	file, err := os.Open("~/code/go/test.go")

	if err != nil {

		// get actual instance which impl error interface
		if err1, ok := err.(*os.PathError); ok {
			fmt.Println(err1.Path)
		}
	} else {

		fmt.Println(file.Name())
	}

	_, err2 := filepath.Glob("[")
	// this == means the error values are same
	if err2 != nil && err2 == filepath.ErrBadPattern {
		fmt.Println(err)
	}

}

// custom define error type by impl error interface
type MyError struct {
	Msg string
}

func (me MyError) Error() string {
	return me.Msg
}

// return custom error pointer by call TestMyError()
func TestMyError() error {
	return &MyError{Msg: "wefwef"}
}

// it's interface define
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
// type error interface {
// 	Error() string
// }
