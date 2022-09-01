package main

import "fmt"

type Employee struct {
	id int
}

type Person struct {
	name string
	age  int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 18
	ftEmployee.id = 1
	GetMessage(ftEmployee.Person)
}
