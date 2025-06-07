package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person = Person{"John", 25}

	fmt.Println("Person Name:", p.Name)
	fmt.Println("Person Age:", p.Age)
}


// This program demonstrates the use of structs in Go.
// It defines a struct type `Person` with fields for name and age,