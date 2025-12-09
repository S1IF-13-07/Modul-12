package main

import "fmt"

func main() {
	var num int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&num)

	if num <= 0 {
		fmt.Println("Input harus bilangan bulat positif.")
		return
	}

	for num > 0 {
		digit := num % 10
		fmt.Println(digit)

		num /= 10
	}
}