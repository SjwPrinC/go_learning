package main

import "fmt"

func main() {
	NestStruct()
	// AnonyStruct()
}

// struct is value type ,so it's default value is {}
func StructDemo() {
	//init 1:
	var p1 Person
	fmt.Println(p1) // print {0}

	//init:2
	p2 := Person{}
	p2.name = "wefwe"
	p2.age = 20
	p2.address = "ShangHai"
	fmt.Println(p2)

	//init:3
	p3 := Person{name: "wefwef", age: 30, address: "Sahanghai"}
	fmt.Println(p3)

	//init:4  new keyword can create point

}

// struct is value type
func StructDemo1() {
	p1 := Person{name: "jerry", age: 10, address: "jiangsu"}
	p2 := p1

	p2.name = "tom"

	fmt.Println(p1, p2)
}

//how to change p1's name? use point
func StructDemo2() {
	p1 := Person{name: "jerry", age: 10, address: "jiangsu"}
	p2 := &p1 //p2 is p1's value

	p2.name = "tom" //omit *, full sytanx is (*p2).name = "tom"

	//*p2 is references p1's value
	fmt.Println(p1, *p2)
}

func StructDemo3() {
	p1 := new(Person) //new can create Person object point
	p1.name = "wefwef"
	p1.age = 100
	p1.address = "zhejiang"

	fmt.Printf("p1 type is %T", p1)
	fmt.Println(*p1)
}

type Person struct {
	name    string
	age     int
	address string
}

func AnonyStruct() {
	s2 := &struct {
		name string
		age  int
	}{
		name: "wef",
		age:  20,
	}

	// fmt.Println(s2)
	// fmt.Printf("%T, ", s2)

	func(obj interface{}) {
		fmt.Println(obj)
	}(s2)
}

//anonymous filed type can't duplicate because same type can't unique set when init
type Worker struct {
	string //anonymous filed
	int
}

func AnonyFiled() {
	w1 := Worker{"worker", 32}

	fmt.Println(w1)
}

//nest struct
type Address struct {
	Province string
	City     string
	Name     string
}

type Doctor struct {
	Address
	Name string
	Age  int
}

func NestStruct() {
	addr := Address{Province: "jiangsu", City: "nantong", Name: "myAddr"}
	d := Doctor{
		Address: addr,
		Name:    "jack",
		Age:     40,
	}

	fmt.Println(d)
	fmt.Println(d.Province)         // if Nest struct's fields are not in field list ,we can omit "Address" and  driect visit Address's province field
	fmt.Println(d.Address.Province) // the two syntax is equal Notice!!! only nest field is valid
	fmt.Println(d.Name)             // can only visit Doctor's Name ,if we need visit Address's name ,we must explicit declare like d.Address.Name
}
