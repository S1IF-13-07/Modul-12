package main
import "fmt"

func main() {
	var token string
	fmt.Print("Masukkan token akses: ")
	fmt.Scan(&token)

	for token != "12345abcde" {
		fmt.Scan(&token)
	}
	fmt.Println("Selamat anda berhasil login")
}