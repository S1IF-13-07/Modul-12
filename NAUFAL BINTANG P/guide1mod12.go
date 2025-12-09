package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    if n == 0 {
        fmt.Println("1")
        return
    }

    for i := n; i >= 1; i-- {
        if i == 1 {
            fmt.Printf("%d", i)
        } else {
            fmt.Printf("%d x ", i)
        }
    }
    fmt.Println()
}