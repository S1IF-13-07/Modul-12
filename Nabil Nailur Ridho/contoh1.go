package main
import "fmt"
func main(){
	var n int
	fmt.Print("Masukkan Bilangan:")
	fmt.Scan(&n)

	for n > 1{
		fmt.Print(n," x ")
		n--
	}
	fmt.Print(1)
}