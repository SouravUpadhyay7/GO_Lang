package main

import "fmt"

func greet(name string) string {
    return "Hello, " + name
}

func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(greet("Alice"))
    fmt.Println("Sum:", add(3, 4))
}
