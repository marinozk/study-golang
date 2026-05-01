package main
import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	
	if x < 0 {
		x = x - (2*x)
		fmt.Print(x)
	} else {
		fmt.Print(x)
	}
}