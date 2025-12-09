package main
import "fmt"
func main() {

	var bil, jumlah int
	fmt.Print("Masukkan sebuah bilangan bulat non negatif : ")
	fmt.Scan(&bil)
	jumlah = bil
	for jumlah > 1 {
		fmt.Print(jumlah, " x ")
		jumlah = jumlah - 1
	}
	fmt.Println(1)
}