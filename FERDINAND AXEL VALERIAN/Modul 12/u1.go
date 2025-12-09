package main

import "fmt"

func main() {
	var gagal int
	var username string
	var password string

	gagal = 0

	for {
		fmt.Print("Masukkan username dan password: ")
		fmt.Scan(&username, &password)

		if username == "Admin" && password == "Admin" {
			break
		} else {
			gagal += 1
			fmt.Println("Username atau Password salah. Silakan coba lagi.")
		}
	}

	fmt.Println(gagal, "percobaan gagal login")
}