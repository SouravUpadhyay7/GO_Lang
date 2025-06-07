package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 3}
	names := []string{"Alice", "Bob"}

	fmt.Println("Numbers:", nums)
	fmt.Println("Names:", names)
}




// This program demonstrates the use of slices in Go.
// It declares a slice of integers and a slice of strings, initializes them,
// and prints their contents to the console.
// Slices are more flexible than arrays and can grow or shrink in size dynamically.