package main

import "fmt"

func main() {
    channel := make(chan int, 3) // Create a buffered channel with capacity 3

    channel <- 10
    channel <- 20
    channel <- 30

    fmt.Println(<-channel) // Output: 10
    fmt.Println(<-channel) // Output: 20
    fmt.Println(<-channel) // Output: 30
}