package main

import "fmt"

func main() {
    var output interface{}
    output = 45

    value, ok := output.(int)
    if ok {
        fmt.Println("Integer value:", value)
    } else {
        fmt.Println("Type assertion failed")
    }
}