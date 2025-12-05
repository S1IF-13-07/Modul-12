package main
import "fmt"

func main() {
    var N, s1, s2, next int
    fmt.Scan(&N)

    s1 = 0
    s2 = 1

    for i := 0; i < N; i++ {
        fmt.Print(s1, " ")
        next = s1 + s2
        s1 = s2
        s2 = next
    }
}
