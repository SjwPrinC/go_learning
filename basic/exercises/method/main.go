package main

import "fmt"

func main() {
	w := Worker{"jerry", 120}
	t := Teacher{"cathy", 100}

	PrintInfo11(w)  // value chuan di
	PrintInfo11(&w) // pointer chuan di
	PrintInfo11(&t) //explicit &t get pointer
	//PrintInfo11(t)  // complie failed

	// pw.PrintInfo()
	// pt.PrintInfo()

	// fmt.Println(w, t)
}

type Worker struct {
	Name string
	Age  int
}

type Teacher struct {
	Name string
	Age  int
}

//impl IInfo
func (w Worker) PrintInfo() {
	w.Age += 10
	fmt.Println(w)
}

//impl IInfo
func (t *Teacher) PrintInfo() {
	t.Age += 10
	fmt.Println(t)
}

type IInfo interface {
	PrintInfo()
}

func PrintInfo11(obj IInfo) {
	obj.PrintInfo()
}
