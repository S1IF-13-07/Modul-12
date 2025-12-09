package main

import (
	"fmt"
)

func main() {
	var username, password string
	jumlahGagal := 0

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		if username == "Admin" && password == "Admin" {
			fmt.Printf("%d percobaan gagal login\n", jumlahGagal)
			break
		} else {
			jumlahGagal++
		}
	}
}
