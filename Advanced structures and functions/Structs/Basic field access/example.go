package main

import "fmt"

// Define the Car struct
type Car struct {
    Name  string
    Type  string
    Brand string
    Years int
}

func main() {
    // Initialize a Car instance
    car := Car{
        Name:  "Model S",
        Type:  "Electric",
        Brand: "Tesla",
        Years: 3,
    }

    // Accessing fields
    fmt.Println("Car Name:", car.Name)
    fmt.Println("Years in Service:", car.Years)
}