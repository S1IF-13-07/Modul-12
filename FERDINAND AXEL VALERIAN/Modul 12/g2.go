package main

import "fmt"

func main() {
	const validToken = "12345abcde"
	var inputToken string

	fmt.Println("Program Login Aplikasi")
	
	for inputToken != validToken {
		fmt.Print("Masukkan token: ")
		fmt.Scan(&inputToken)

		if inputToken == validToken {
			fmt.Println("Selamat Anda berhasil login")
		} else {
			fmt.Println("Token salah. Silakan coba lagi.")
		}
	}
}