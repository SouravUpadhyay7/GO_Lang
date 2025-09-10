package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from Goroutine!")
}

func main() {
    go sayHello() // Starts a new goroutine
    time.Sleep(1 * time.Second) // Allow goroutine to complete
    fmt.Println("Main function finished")
}