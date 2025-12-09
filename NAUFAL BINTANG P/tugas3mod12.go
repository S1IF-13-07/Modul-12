package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan dua bilangan: ")
	fmt.Scan(&a, &b)

	hasil := 0
	for a >= b { 
		a = a - b
		hasil++
	}

	fmt.Println(hasil)
}