package main
import "fmt"

func main() {
	var x float64
	
	fmt.Scanln(&x)
	
	fmt.Println((x*x*x + 3*x) / (x*x*x*x - 3*x*x +4))
}