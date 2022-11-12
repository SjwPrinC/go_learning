package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mi MyInt = 100
	var ms MyString = "jerry"

	var i int
	var s MyString1

	i = int(mi) // type MyInt is unmatch int,so we need use i = int(mi)
	s = ms      //MyString1  is equals  MyString and string  it's funny

	fmt.Println(i, s)

	fmt.Printf("type of func is : %T \n", func1())
	fmt.Println(func1()(1, 2))
}

// type redefine  type MyInt is unmatch int
// redefine type can extern type ,like following example:
type MyInt int

// extern Convert2Str method by using MyInt type
func (mi MyInt) Convert2Str() string {
	return strconv.Itoa(int(mi))
}

// type alias    it means MyString is equals string and they can switch each other
type MyString = string
type MyString1 = string

// type also can define func type
type MyFunc = func(int, int) string

// redefine func(int, int) string as MyFunc1
type MyFunc1 = func(int, int) string

// use func alias as return value
func func1() MyFunc {
	return func(int, int) string {
		return "wef"
	}
}
