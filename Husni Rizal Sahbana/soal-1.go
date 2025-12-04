package main

import "fmt"

func main() {
	const Username = "Admin"
	const Password = "Admin"

	var username, password string
	var gagal int

	for {
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&username)

		fmt.Print("Masukkan Password: ")
		fmt.Scan(&password)

		if username == Username && password == Password {
			break
		} else {
			gagal++
			fmt.Println("Username atau password salah! Coba lagi.\n")
		}
	}

	fmt.Printf("%d percobaan gagal login\n", gagal)
}
