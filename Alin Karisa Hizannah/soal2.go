package main

import "fmt"

func main() {
    var num,digit int
    fmt.Print("Masukkan sebuah bilangan bulat positif: ")
    fmt.Scan(&num)

    for num > 0 {
        digit = num % 10
        fmt.Println(digit)
        num = num / 10
    }
}
