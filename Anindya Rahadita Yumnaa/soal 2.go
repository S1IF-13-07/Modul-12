package main

import "fmt"

func main() {
    var n int
	fmt.Print("Masukkan angka: ")
    fmt.Scan(&n)

    for n > 0 {
        fmt.Println(n % 10)
        n = n / 10
	}
}