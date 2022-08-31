package main

import "fmt"

// Struct
type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)

	e.id = 1
	e.name = "Name"
	fmt.Printf("%v", e)
}
