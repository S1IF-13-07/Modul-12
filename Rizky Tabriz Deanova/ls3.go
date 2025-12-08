package main

import (
	"fmt"
)

func main() {
	var x, y int
	var hasil int
	hasil = 0
	fmt.Scan(&x, &y)
	for x >= y {
		x = x - y
		hasil++
	}
	fmt.Println(hasil)
}
