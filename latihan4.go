package main
import "fmt"

func main() {
	var x, y float64
	
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	
	fmt.Println((1/(3*x*x + 10)) +10*y + 7)
}