package main
import "fmt"

func main() {
	var x, i, d1, d2, hasil int
	fmt.Scan(&x)
	d1 = 0
	d2 = 1
	fmt.Print(d1, " ", d2)
	
	for i = 2; i <= x; i++ {
		hasil = d1 + d2
		fmt.Print(" ", hasil)
		d1 = d2
		d2 = hasil
	}
}