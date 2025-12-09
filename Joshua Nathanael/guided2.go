package main

import "fmt"

func main() {
	var token string
	for token != "12345abcde" {
		fmt.Print("MASUKAN TOKEN : ")
		fmt.Scan(&token)
	}
	fmt.Println("KAMU BERHASIL LOGIN")
}