package main

import "fmt"

func main() {
	A()
}

func A() {
	i := 100

	// the function exection has been delayed, so at this line i is 2 and send to func so println(a) is 2
	defer func(a int) {
		fmt.Println(a)
	}(i)

	i++

	// panic("error has hanppend")

	defer func(a int) {
		fmt.Println(a)
	}(i)

}
