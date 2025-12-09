package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangannya massss: ")
    fmt.Scan(&n)

    fmt.Print("Keluaran: ")

    if n == 0 {
        fmt.Print("1") 
    } else {
        for i := n; i >= 1; i-- {
            if i == 1 {
                fmt.Printf("%d", i)
            } else {
                fmt.Printf("%d x ", i)
            }
        }
    }

    fmt.Println()
}