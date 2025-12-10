package main
import "fmt"
func main() {
	var a int
	fmt.Print("Masukkan sebuah bilangan untuk di jadikan fibonanci : ")
	fmt.Scan(&a)
	n1, n2 := 0, 1
	fmt.Printf("Deret Fibonacci hingga %d adalah: \n", a)
	for n1 <= a {
		fmt.Printf("%d ", n1)
		n1, n2 = n2, n1+n2
	}
	fmt.Println()


}