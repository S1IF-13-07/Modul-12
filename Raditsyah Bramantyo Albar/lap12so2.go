package main

import "fmt"

func main() {
	var n, urut int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	for n > 0 {
		urut = n % 10
		fmt.Println(urut)
		n = n / 10
	}
}
