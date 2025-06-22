package main

import "fmt"

func main() {
    i := 0
START:
    if i < 3 {
        fmt.Println("i:", i)
        i++
        goto START
    }
}
