package main
import "fmt"

func main() {
	var a string
	var b string
	fmt.Print("Masukkan username : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan password : ")
	fmt.Scan(&b)

	for a != "Admin" || b != "Admin" {
		fmt.Println("4 Percobaan Gagal Login")
		fmt.Print("Masukkan username : ")
		fmt.Scan(&a)
		fmt.Print("Masukkan password : ")
		fmt.Scan(&b)
	}
	fmt.Println("0 Percobaan Gagal Login")
}