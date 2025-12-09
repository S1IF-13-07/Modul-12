package main
import "fmt"
func main() {
	var a int

	fmt.Print("Masukan bilangan bulat positif : ")
	fmt.Scan(&a)

	for a > 0 {
		fmt.Println(a % 10)
		a /= 10
	}
}