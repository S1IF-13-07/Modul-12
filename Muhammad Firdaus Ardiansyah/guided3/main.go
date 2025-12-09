package main

import "fmt"

func main() {
	var a, s1, s2, b, c int
	fmt.Scan(&a)
	s1 = 0
	s2 = 1
	b = 0
	for b < a {
		fmt.Print(s1, " ")
		c = s1 + s2
		s1 = s2
		s2 = c
		b = b + 1
	}
}