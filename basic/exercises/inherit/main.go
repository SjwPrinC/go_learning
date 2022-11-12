package main

import "fmt"

func main() {
	s := Student{
		Person: Person{"jerry", "male", 20},
		School: "ShangHai School",
		Class:  "Gao yi",
	}

	s.Eat()   // if Student is nest person, we can call person's methods omit "Person"
	s.Study() // if Student and Person both have Study() method, should explicit call,the fllowing
	s.Person.Study()
}

type Person struct {
	Name string
	Sex  string
	Age  int
}

type Student struct {
	Person
	School string
	Class  string
}

func (p Person) Eat() {
	fmt.Printf(" person %s is eatting \n", p.Name)
}

func (s Student) Study() {
	fmt.Printf("Student %s is studying \n", s.Name)
}

func (p Person) Study() {
	fmt.Printf(" person %s is eatting \n", p.Name)
}
