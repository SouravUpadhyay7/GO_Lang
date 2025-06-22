package main

import "fmt"

func main() {
    // Loop
    for i := 1; i <= 5; i++ {
        fmt.Println("i:", i)
    }

    // If-Else
    num := 7
    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // Switch
    day := 2
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    default:
        fmt.Println("Other day")
    }
}
