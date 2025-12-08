package main

import "fmt"

func main() {
	var x, y, count int

	fmt.Scan(&x)
	fmt.Scan(&y)

	for x >= y {
		count++
		x -= y
	}
	fmt.Print(count)
}
