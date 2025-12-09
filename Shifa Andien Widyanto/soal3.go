package main
import "fmt"
func main() {
	var x, y int

	fmt.Print("masukan dua bilangan bulat positif :")
	fmt.Scan(&x, &y)
	hasil := 0

	for x >= y {
		x -= y
		hasil++
	}
	fmt.Print(hasil)
}