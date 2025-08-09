package main
import "fmt"

func main() {
	var name, pass string
	var gagal int = 0
	
	for fmt.Scan(&name, &pass); name != "admin" || pass != "admin"; fmt.Scan(&name, &pass) {
		gagal++
	}
	fmt.Println(gagal)
	fmt.Println("Login berhasil")
}
	