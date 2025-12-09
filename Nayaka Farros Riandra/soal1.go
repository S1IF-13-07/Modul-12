package main

import "fmt"

func main(){
	var user, pass string
	var attempt int
	attempt = 0
	
	for {
		fmt.Print("Masukkan username : ")
		fmt.Scan(&user)
		fmt.Print("Masukkan password : ")
		fmt.Scan(&pass)

		if user == "Admin" && pass == "Admin"{
			break
			} else {
				attempt += 1
				fmt.Println("Username atau password salah, silahkan coba lagi")
			}
		}
		fmt.Println(attempt , "percobaan gagal login")
}