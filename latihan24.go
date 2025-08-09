package main
import "fmt"

func main() {
	var a, hasil int
	
	for fmt.Scan(&a); a % 2 == 0; fmt.Scan(&a) {
		hasil = hasil + a
	}
	fmt.Println(hasil)
}