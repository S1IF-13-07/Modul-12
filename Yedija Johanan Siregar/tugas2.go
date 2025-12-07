package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)

  if n == 0 {
    fmt.Println(0)
  } else {
    for n > 0 {
    digit := n % 10
      fmt.Println(digit)
      n = n / 10
    }  
  }
}
