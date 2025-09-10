package main

import "fmt"

func sendData(channel chan int) {
    channel <- 42 // Send data into the channel
}

func main() {
    channel := make(chan int) // Create an unbuffered channel
    go sendData(channel)      // Start goroutine to send data
    fmt.Println(<-channel)    // Receive data from the channel
}