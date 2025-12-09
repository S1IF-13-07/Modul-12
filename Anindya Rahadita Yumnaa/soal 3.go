package main

import "fmt"

func main() {
    var x, y int
	fmt.Print("Masukkan dua angka: ")
    fmt.Scan(&x, &y)

    hasil := 0

    for x >= y {
        x = x - y
        hasil++
    }

    fmt.Println(hasil)
}