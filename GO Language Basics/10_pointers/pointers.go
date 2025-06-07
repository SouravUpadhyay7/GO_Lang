package main

import "fmt"

func main() {
	var p *int
	x := 42
	p = &x

	fmt.Println("Value of x:", *p)
}


// This program demonstrates the use of pointers in Go.
// It declares a pointer variable `p`, assigns it the address of an integer variable `x`,
// and then prints the value stored at that address using dereferencing.
// Pointers are used to reference memory locations directly, allowing for efficient data manipulation.