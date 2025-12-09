package main

import "fmt"

func main() {
	var token string
	fmt.Scan(&token)

	for token != "1234abcde" {
		fmt.Scan(&token)
	}
	fmt.Print("Selamat anda berhasil ")
}
