package main

import "fmt"

func main() {
	var user, pass string
	user1 := "Admin"
	pass1 := "Admin"
	i := 0
	for {
		fmt.Scan(&user, &pass)
		if user == user1 && pass == pass1 {
			fmt.Printf("%v percobaan gagal login", i)
			break
		} else {
			i++
		}
	}
}
