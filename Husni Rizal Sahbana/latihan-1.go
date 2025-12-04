package main

import "fmt"

func main() {

	var bilangan, j int
	fmt.Print("Masukan Bilangan : ")
	fmt.Scan(&bilangan)

	j = bilangan
	for j > 1 {
		fmt.Print(j, " x ")
		j--
	}

	fmt.Println(1)
}
