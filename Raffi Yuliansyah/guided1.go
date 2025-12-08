package main

import "fmt"

func main(){
	var num int
	fmt.Scan(&num)

	for num > 1 {
		fmt.Print(num, " x ")
		num--
	}
	fmt.Print(1)
}