package main

import . "fmt"

// import _ "os.File"

func main() {
	Println("hello")
}

type Person struct {
	Name string // if first letter is upper, this field will be visited by other package
	age  int    // age field only visited by the same package
}

// init function will be called when init package
// init function can be duplicate declare
func init() {
	Println("first init function been called")
}

func init() {
	Println("second init function been called")
}

// 1. import sequence is the sequence of init() func exection

// the exec sequence of main package 1. const 2. var .... 3.init() 4.mian
