package main

import "fmt"

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting risky operation...")
	panic("something went wrong!") // trigger panic
	fmt.Println("This line will not run")
}

func main() {
	riskyOperation()
	fmt.Println("Program continues after recovery")
}
