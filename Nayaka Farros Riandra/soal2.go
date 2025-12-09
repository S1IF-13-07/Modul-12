package main 
import "fmt"

func main() {
	var angka, urut int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	for angka > 0 {
		urut = angka % 10
		fmt.Println(urut)
		angka = angka / 10
	}
}