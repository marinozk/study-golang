package main
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	var hasil int
	
	for a > 0 {
		var mod int = a % 10
		fmt.Print(" ", mod)
		hasil = hasil + mod
		a = a / 10
	}
	fmt.Println("\n", hasil)
}