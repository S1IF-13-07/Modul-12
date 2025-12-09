package main
import "fmt"
func main() {
	var n, a int

	fmt.Print("Masukan bilangan bulat positif n :")
	fmt.Scan(&n)
	a = n 
	for a > 1 {
		fmt.Print(a, " x ")
		a = a - 1
	}
	fmt.Println(1)	
	} 
	
