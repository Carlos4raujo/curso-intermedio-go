package main

import "fmt"

// Struct
type Employee struct {
	id   int
	name string
}

// Methods
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e Employee) GetId() int {
	return e.id
}

func (e Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	// fmt.Printf("%v", e)

	e.id = 1
	e.name = "Name"
	// fmt.Printf("%v", e)

	e.SetId(2)
	e.SetName("Hello")
	// fmt.Printf("%v", e)

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
