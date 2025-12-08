package main

import (
	"fmt"
)

func main() {
	var user, pass string
	percobaan := 0

	for {
		fmt.Print("Username: ")
		fmt.Scan(&user)
		fmt.Print("Password: ")
		fmt.Scan(&pass)

		if user == "Admin" && pass == "Admin" {
			break
		}

		percobaan++
		fmt.Printf("Username atau password salah, coba lagi.\n")
	}

	fmt.Printf("%d percobaan gagal login\n", percobaan)
}
