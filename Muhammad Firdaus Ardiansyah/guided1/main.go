package main

import "fmt"

func main() {
	var n int 
	fmt.Print("Masukkan bilangan bulat non negatif: ")
	fmt.Scan(&n)
	for n > 1 {
		fmt.Print(n, " x ")
		n--
	}
	fmt.Println(1)
}