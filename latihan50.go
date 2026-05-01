package main
import "fmt"

func main() {
	var a, b, c int
	var hasil bool
	fmt.Scan(&a, &b, &c)
	
	hasil = (a + b == 2*c) || (a + c == 2*b) || (b + c == 2*a)
	fmt.Print(hasil)
}