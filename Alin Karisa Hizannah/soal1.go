package main

import "fmt"

func main() {
    var usr, pwd string
    var count int

    for {
        fmt.Scan(&usr, &pwd)

        if usr == "Admin" && pwd == "Admin" {
            fmt.Printf("%d percobaan gagal login\n", count)
            break
        } else {
            count++
        }
    }
}
