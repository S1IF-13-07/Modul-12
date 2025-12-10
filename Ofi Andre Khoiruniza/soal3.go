package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	hasil := 0
	total := y

	for total <= x {
		hasil++
		total += y
	}

	fmt.Println(hasil)
}
