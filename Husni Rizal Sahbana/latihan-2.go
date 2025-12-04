package main

import "fmt"

func main() {

	var token string
	fmt.Print("Masukan Token : ")
	fmt.Scan(&token)

	for token != "12345abcde" {
		fmt.Print("Masukan Token : ")
		fmt.Scan(&token)
	}

	fmt.Println("Selamat Token Anda Benar")
}