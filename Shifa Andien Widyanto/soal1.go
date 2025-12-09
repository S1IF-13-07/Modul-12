package main
import "fmt"
func main() {
	var u, p string
	gagal := 0

	for{
		fmt.Scan(&u, &p)
		if u == "Admin" && p == "Admin" {
			break
		}
		gagal++
	}
	fmt.Printf("%d percobaan gagal login\n", gagal)
}