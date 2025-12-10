package main
import "fmt"

func main() {
	var a int
	fmt.Print("masukkan sebuah angka untuk di jadikan factorial : ")
	fmt.Scan(&a)
	

	for a > 1 {
		fmt.Print(a, "x")
		a--
		
	}
	fmt.Println(1)

}