package main
import "fmt"

func main() {
	var x, hasil, d1, d2, d3 int
	var diskon, cashback bool
	fmt.Scanln(&x)
	hasil = (x / 10) % 100
	diskon = hasil % 2 == 0
	d1 = x % 10
	d2 = (x / 10) % 10
	d3 = x / 1000
	cashback = (d1 != 0) && ((d3 + d2) % d1 == 0)
	
	fmt.Println("diskon?", diskon)
	fmt.Println("Cashback", cashback)
	
}