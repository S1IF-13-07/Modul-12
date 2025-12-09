package main

import "fmt"

func main() {
	const Username = "Admin"
	const Password = "Admin"

	var username, password string
	var gagal int

	for {
		fmt.Print("MASUKAN USERNAME: ")
		fmt.Scan(&username)

		fmt.Print("MASUKAN PASSWORD: ")
		fmt.Scan(&password)

		if username == Username && password == Password {
			break
		} else {
			gagal++
			fmt.Println("USERNAME ATAU PASSWORD SALAH! COBA LAGI.\n")
		}
	}

	fmt.Printf("%d PERCOBAAN GAGAL LOGIN\n", gagal)
}