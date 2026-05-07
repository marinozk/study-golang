package main
import "fmt"

func funcLuas(r float64, t float64) (hasil float64) {
	hasil = 2 * 3.14 * r  * (r + t)
	return hasil
}

func main() {
	var r, t float64
	fmt.Scan(&r, &t)
	
	hasil := funcLuas(r, t)
	fmt.Printf("%.3f\n", hasil)
}