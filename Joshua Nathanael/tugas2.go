package main

import "fmt"

func main() {
	var num int 
	fmt.Print("MASUKAN ANGKANYA MASSS: ")
	fmt.Scan(&num)

	for num > 0 {
		digit := num % 10
		fmt.Println(digit)
		num /= 10
	}
}