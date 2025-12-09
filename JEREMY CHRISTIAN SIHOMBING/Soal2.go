package main 
import "fmt"
func main() {

	var nilai, berurut int
	fmt.Print("Masukkan suatu bilangan bulat positif: ")
	fmt.Scan(&nilai)
	for nilai > 0 {
		berurut = nilai % 10
		fmt.Println(berurut)
		nilai = nilai / 10
	}
}