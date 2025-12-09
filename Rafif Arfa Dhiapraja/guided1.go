package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan sebuah Bilangan Non Negatif: ")
	fmt.Scan(&n)

	for n > 1 {
		fmt.Print(n, " x ")
		n--
	}

	fmt.Print(1)

}
