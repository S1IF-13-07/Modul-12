package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	i := 1

	if i < n {
		f1 := 0
		f2 := 1
		fmt.Print(f1, " ", f2, " ")
		n -= 2
		for i <= n {
			next := f2
			f2 += f1
			fmt.Print(f2, " ")
			f1 = next
			i++
		}
	} else {
		fmt.Print("Masukan angka yang benar")
	}
}
