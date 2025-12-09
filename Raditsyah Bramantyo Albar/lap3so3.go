package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&x, &y)

	hasil = 0

	for x >= y {
		x = x - y
		hasil++
	}
	fmt.Print(hasil)
}
