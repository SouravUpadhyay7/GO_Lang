package main

import "fmt"

var globalVar = "I am global"

func showScope() {
    localVar := "I am local"
    fmt.Println(globalVar)
    fmt.Println(localVar)
}

func main() {
    showScope()
    fmt.Println(globalVar)
    // fmt.Println(localVar) // Error: undefined
}
