package main

import "fmt"

func main() {
	c := &Cat{"jerry"}

	var cc A = c
	cc.test1()

	fmt.Printf("cc type is %T\n", cc)

	var ccc C = c
	ccc.test3()
	fmt.Printf("ccc type is %T\n", ccc) // the type of ccc and cc both main.Cat

	// type assert   instance :=interfaceObj.(real Type) it will be panic
	// instance,ok :=interfaceObj.(real Type) it's safe when type unmatch ok is false
	if instance, ok := ccc.(*Cat); ok {
		fmt.Println(instance)
	}

	switch ins := ccc.(type) {
	case Cat:
		fmt.Println("my type is Cat, value is :", ins)

	case *Cat:
		fmt.Println("my type is *Cat")
	default:
		fmt.Println("unknown type")
	}
}

type A interface {
	test1()
}

type B interface {
	test2()
}

//nest interface like nest struct
// if you  want impl interface C ,you must impl interface A  and B's method
type C interface {
	A
	B
	test3()
}

type Cat struct {
	string
}

func (c Cat) test1() {
	fmt.Println("test1()")
}

func (c Cat) test3() {
	fmt.Println("test3()")
}

func (c Cat) test2() {
	fmt.Println("test3()")
}
