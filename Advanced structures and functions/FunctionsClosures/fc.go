package main

import "fmt"

// A function that returns another function (closure)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// Create two independent counters
	counter1 := counter()
	counter2 := counter()

	fmt.Println("Counter 1:", counter1()) // 1
	fmt.Println("Counter 1:", counter1()) // 2
	fmt.Println("Counter 1:", counter1()) // 3

	fmt.Println("Counter 2:", counter2()) // 1
	fmt.Println("Counter 2:", counter2()) // 2
}
