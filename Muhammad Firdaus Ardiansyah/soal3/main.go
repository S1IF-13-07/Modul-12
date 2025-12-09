package main 
import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan dua angka: ")
	fmt.Scan(&x, &y)
	count := 0
	
	for x >= y {
		x -= y
		count++
	}
	fmt.Println(count)
}