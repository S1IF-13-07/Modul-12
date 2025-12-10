package main
import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Masukkan Angka Pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan Angka Kedua : ")
	fmt.Scan(&b)
	
	if b <= 0 {
		fmt.Println("Bilangan b harus lebih dari 0.")
		return
	}
	hasil := 0
	total := b

	for total <= a {
		hasil++
		total += b
	}

	fmt.Println(hasil)

}