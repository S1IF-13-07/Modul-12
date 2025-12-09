package main

import "fmt"

func main() {
	var username, password string
	fmt.Print("Masukkan Username & Password: ")
	fmt.Scan(&username, &password)

	jumlah := 0

	for username != "Admin" || password != "Admin" {
		jumlah = jumlah + 1
		fmt.Print("Username & Password Salah! Masukkan Kembali: ")
		fmt.Scan(&username, &password)
	}

	fmt.Printf("%d Percobaan gagal", jumlah)
}