package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&bilangan)

	fmt.Println("Digit bilangan dari Kanan ke Kiri: ")

	for bilangan != 0 {
		digit := bilangan % 10
		fmt.Println(digit)
		bilangan = bilangan / 10
	}
}