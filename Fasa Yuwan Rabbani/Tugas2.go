package main
import "fmt"

func main() {
	var a int
	fmt.Print("Masukkan Sebuah Angka Bilangan Bulat : ")
	fmt.Scan(&a)

	for a > 0 {
		digit := a % 10
		fmt.Println(digit)
		a = a / 10
	} 

}