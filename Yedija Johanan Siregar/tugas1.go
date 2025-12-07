package main

import "fmt"

func main() {
  var username, password string
  count := 0
  for {
      fmt.Scan(&username, &password)
      if username == "Admin" && password == "Admin" {
    
        break
    }
      count++
  }
    fmt.Printf("%d percobaan gagal login", count)
}
