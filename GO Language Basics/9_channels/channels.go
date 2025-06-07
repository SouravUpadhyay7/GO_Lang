package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from channel!"
	}()

	msg := <-ch
	fmt.Println(msg)
}
// This program creates a channel, sends a message from a goroutine, and receives it in the main function.