package main
import "fmt"
func main (){
	var usn, pw string
	gagal := 0
	fmt.Print("Masukkan username & password: ")
	fmt.Scan(&usn, &pw)

	for usn != "Admin" || pw != "Admin"{
		gagal++
		fmt.Scan(&usn, &pw)
	}
	fmt.Printf("%d percobaan gagal login\n", gagal)
}