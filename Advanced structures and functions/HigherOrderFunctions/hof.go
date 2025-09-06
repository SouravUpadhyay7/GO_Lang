package main

import "fmt"

// A higher-order function that takes another function as argument
func operate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Example functions
func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func main() {
	// Using higher-order function with add
	sum := operate(10, 5, add)
	fmt.Println("Addition:", sum)

	// Using higher-order function with multiply
	product := operate(10, 5, multiply)
	fmt.Println("Multiplication:", product)

	// Using an anonymous function directly
	diff := operate(10, 5, func(x int, y int) int {
		return x - y
	})
	fmt.Println("Subtraction:", diff)
}
