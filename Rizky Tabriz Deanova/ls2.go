package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)
	for bil > 0 {
		bil1 := bil % 10
		fmt.Println(bil1)
		bil = bil / 10
	}
}
