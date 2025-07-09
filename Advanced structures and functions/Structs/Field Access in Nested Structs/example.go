package main

import "fmt"

// Define the Engine struct
type Engine struct {
    HorsePower int
    Mileage    float64
}

// Define the Car struct
type Car struct {
    Name   string
    Engine Engine
}

func main() {
    // Initialize a Car instance
    car := Car{
        Name: "Model X",
        Engine: Engine{
            HorsePower: 300,
            Mileage:    20.5,
        },
    }

    // Accessing fields
    fmt.Println("Car Name:", car.Name)
    fmt.Println("HorsePower:", car.Engine.HorsePower)
    fmt.Println("Mileage:", car.Engine.Mileage)
}