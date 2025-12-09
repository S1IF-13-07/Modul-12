package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	for angka > 0 {
		digit := angka % 10 
		fmt.Println(digit)  
		angka = angka / 10 
	}
}