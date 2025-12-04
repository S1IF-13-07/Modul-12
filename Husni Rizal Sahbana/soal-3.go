package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	count := 0

	for x >= y {
		x = x - y
		count++
	}

	fmt.Println(count)
}
