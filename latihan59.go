package main
import "fmt"

func funcVolume(r float64, t float64) (hasil float64) {
	hasil = 3.14 * (r*r) * t
	return hasil
}

func main() {
	var r, t float64
	fmt.Scan(&r, &t)
	
	hasil := funcVolume(r, t)
	fmt.Printf("%.1f\n", hasil)
}