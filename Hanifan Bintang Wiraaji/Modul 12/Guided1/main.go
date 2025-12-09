package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Print("1")
	} else if n > 0 {
		for n > 0 {
			fmt.Print(n)
			if n == 1 {
				break
			}
			fmt.Print(" x ")
			n--
		}
	} else {
		fmt.Print("Masukan bilangan positif")
	}
}
