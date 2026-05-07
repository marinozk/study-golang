package main
import "fmt"

func konversi(suhu float64) float64 {
	return (1.8 * suhu) + 32
}

func main() {
	var n int
	var suhu float64
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&suhu)
		hasil := konversi(suhu)
		fmt.Println(suhu, " Celsius = ", hasil, " Fahrenheit")
	}
	
}