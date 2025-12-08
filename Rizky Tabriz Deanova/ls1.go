package main

import (
	"fmt"
)

func main() {
	var username, password string
	var gagal int
	gagal = 0
	for {
		fmt.Scan(&username, &password)
		if username == "Admin" && password == "Admin" {
			break
		} else {
			gagal++
		}
	}
	fmt.Printf("%d percobaan gagal login\n", gagal)
}
