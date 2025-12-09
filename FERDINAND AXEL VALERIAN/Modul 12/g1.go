package main

import "fmt"

func main() {
	var num int

	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Input harus bilangan bulat non-negatif.")
		return
	}

	if num == 0 || num == 1 {
		fmt.Println(num, " -> 1")
		return
	}

	fmt.Printf("%d -> ", num)
	
	for i := num; i > 1; i-- {
		fmt.Printf("%d x ", i)
	}

	fmt.Println("1")
}