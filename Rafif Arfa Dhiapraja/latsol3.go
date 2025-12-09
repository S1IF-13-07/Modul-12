package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan dua bilangan bulat: ")
	fmt.Scan(&x, &y)

	jumlah := 0

	for x >= y {
		x = x - y
		jumlah = jumlah + 1
	}

	fmt.Println("Hasil dari pembagian adalah:", jumlah)
}