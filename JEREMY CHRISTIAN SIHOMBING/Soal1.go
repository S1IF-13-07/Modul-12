package main 
import "fmt"
func main() {

	var gagal int
	var name, pw string
	gagal = 0
	for {
		fmt.Print("Masukkan Username dan Password: ")
		fmt.Scan(&name, &pw)
		if name == "Admin" && pw == "Admin"{
			break
			} else {
				gagal += 1
			}
		}
		fmt.Println(gagal , "percobaan gagal login")
}