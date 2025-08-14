package main
import "fmt"

func main() {
	var a int 
	fmt.Scanln(&a)
	var hasil int
	
	for a > 0 {
		var digit int = a % 10
		fmt.Print(digit, "")
		hasil = hasil + digit
		a = a / 10
	}
	fmt.Println()
	fmt.Println(hasil)
}