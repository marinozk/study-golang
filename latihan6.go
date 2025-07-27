package main
import "fmt"

func main() {
	var x, d1, sisa, d2, d3 int
	
	fmt.Scanln(&x)
	
	d1 = x / 100
	sisa = x % 100
	d2 = sisa / 10
	sisa = sisa % 10
	d3 = sisa / 1
	
	
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
} 