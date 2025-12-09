package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("MASUKAN X DAN Y: ")
	fmt.Scan(&x, &y)

	hasil := 0
	sisa := x
	
	for sisa >= y {
		sisa -= y   
		hasil++     
	}

	fmt.Println("HASIL:", hasil)
}