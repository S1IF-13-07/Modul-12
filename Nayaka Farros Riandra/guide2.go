package main

import "fmt"

func main() {
	var token string
	for token != "12345abcde" {
		fmt.Print("masukkan token : ")
		fmt.Scan(&token)
	}
	fmt.Println("anjayy berhasil login")
}