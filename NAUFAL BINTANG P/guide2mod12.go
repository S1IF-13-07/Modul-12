package main

import "fmt"

func main() {
    const validToken = "12345abcde"
    var token string

    for {
        fmt.Print("Masukkan token: ")
        fmt.Scanln(&token)

        if token == validToken {
            fmt.Println("Selamat Anda berhasil login")
            break
        }
    }
}