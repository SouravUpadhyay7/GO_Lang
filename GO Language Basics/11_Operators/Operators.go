package main

import "fmt"

func main() {
    a, b := 10, 3

    fmt.Println("Addition:", a + b)
    fmt.Println("Subtraction:", a - b)
    fmt.Println("Multiplication:", a * b)
    fmt.Println("Division:", a / b)
    fmt.Println("Modulus:", a % b)

    fmt.Println("a == b:", a == b)
    fmt.Println("a != b:", a != b)
    fmt.Println("a > b:", a > b)
    fmt.Println("a < b:", a < b)

    fmt.Println("Logical AND:", a > 5 && b < 5)
    fmt.Println("Logical OR:", a < 5 || b < 5)
}
