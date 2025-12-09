package main

import "fmt"

func main() {
    const correctUser = "Admin"
    const correctPass = "Admin"

    var user, pass string
    gagal := 0

    for {
        fmt.Scan(&user, &pass)

        if user == correctUser && pass == correctPass {
            break
        }

        gagal++
    }

    fmt.Printf("%d percobaan gagal login\n", gagal)
}