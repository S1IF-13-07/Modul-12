package main

import "fmt"

func main() {
	var inputPass, pass, inputUser, user string
	var count int
	user = "Admin"
	pass = "Admin"
	count = 0
	fmt.Scan(&inputUser)
	fmt.Scan(&inputPass)
	for inputUser != user || inputPass != pass {
		fmt.Scan(&inputUser)
		fmt.Scan(&inputPass)
		count++
	}
	fmt.Printf("%d percobaan gagal login", count)
}
