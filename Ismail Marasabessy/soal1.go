package main

import "fmt"

func main() {
	const userBenar = "Admin"
	const passBenar = "Admin"

	var user, pass string
	gagal := 0

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&user)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pass)

		if user == userBenar && pass == passBenar {
			break
		}

		gagal++
		fmt.Println("Username atau password salah, coba lagi.")
	}

	fmt.Printf("%d percobaan gagal login\n", gagal)
}
