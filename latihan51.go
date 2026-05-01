package main
import "fmt"

func main() {
	var r1, r2, pusat int
	var hasil bool
	fmt.Scan(&r1, &r2, &pusat)
	
	hasil = r1 + r2 < pusat == true || r1 + r2 >= pusat == false
	fmt.Print(hasil)
}