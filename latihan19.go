package main
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	
	var total float32 = 0
	
	for i := 0; i < a; i++ {
		var d float32
		fmt.Scanln(&d)
		total = total + d
	}
	var hasil float32
	hasil = total / float32(a)
	fmt.Printf("%.3f\n", hasil)
}