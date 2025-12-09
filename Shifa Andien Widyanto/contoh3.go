package main
import "fmt"
func main() {
	var n, a1,  a2, j, temp int

	fmt.Print("Masukan bilangan bulat positif : ")
	fmt.Scan(&n)
	a1 = 0
	a2 = 1
	j = 0

	for j < n {
		fmt.Print(a1, " ")
		temp = a1 + a2
		a1 = a2
		a2 = temp
		j = j + 1
	}
}