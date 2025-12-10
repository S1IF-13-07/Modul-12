package main

import "fmt"

func main() {
	var user, pass string
	gagal := 0

	fmt.Scan(&user, &pass)
	for user != "Admin" || pass != "Admin" {
		gagal++
		fmt.Scan(&user, &pass)
	}

	fmt.Println(gagal, "percobaan gagal login")
}
