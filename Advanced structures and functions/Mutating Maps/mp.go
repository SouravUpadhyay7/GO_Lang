package main

import "fmt"

func main() {
	// Creating a map
	students := map[string]int{
		"Alice": 85,
		"Bob":   90,
		"Carol": 92,
	}

	fmt.Println("Original Map:", students)

	// Mutating (updating) a value
	students["Alice"] = 95
	fmt.Println("After updating Alice's marks:", students)

	// Adding a new key-value pair
	students["David"] = 88
	fmt.Println("After adding David:", students)

	// Deleting a key-value pair
	delete(students, "Bob")
	fmt.Println("After deleting Bob:", students)

	// Accessing a key safely
	if val, exists := students["Carol"]; exists {
		fmt.Println("Carol's marks:", val)
	} else {
		fmt.Println("Carol not found")
	}
}
