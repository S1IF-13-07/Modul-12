package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Scan(&x, &y)
	bagi := y
	for bagi <= x {
		hasil += 1
		bagi += y
	}
	fmt.Print(hasil)
}
