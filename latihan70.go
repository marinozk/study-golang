package main
import "fmt"

func konversi(suhu float64, r, f, k *float64) {
	*r = 0.8 * suhu
	*f = 1.8 * suhu + 32
	*k = suhu + 273.15
}

func main() {
	var suhu float64
	var r, f, k float64
	fmt.Scan(&suhu)
	
	konversi(suhu, &r, &f, &k)
	fmt.Print(r, "R ", f, "F ", k, "K ")
	
}