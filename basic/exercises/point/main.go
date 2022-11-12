package main

import "fmt"

func main() {

	PointArrayTest()

	// var a int = 10

	// var ap *int

	// ap = &a

	// //*ap is to retrive value in address ap
	// *ap = 300
	// fmt.Println("a is : ", a)
	// fmt.Printf("a address is %p \n", &a)
	// fmt.Println("a point address is: ", ap)

	// app := &ap

	// fmt.Printf("a type is %T, ap type is %T app type is %T \n", a, ap, app)
}

func ArrayPointTest() {
	arr1 := [4]int{1, 2, 3, 4}

	arr1p := &arr1

	fmt.Println(arr1, arr1p)

	fmt.Println((*arr1p)[0])

	(*arr1p)[0] = 222 //normal sytanx is *arr1p to get array value
	arr1p[0] = 222    //but we can omit * ,and set value 222 to arr1p[0]

	fmt.Println(arr1, arr1p)
}

//the array element type is *type
func PointArrayTest() {
	a, b, c, d := 1, 2, 3, 4
	pa := []*int{&a, &b, &c, &d}

	fmt.Printf("pa type is %T \n", pa)
	fmt.Println(pa)

	*pa[0] = 1233
	fmt.Println(*pa[0]) //get value of first element
}
