package main
import "fmt"

func main() {
	var a, b, c, d, e float32 
	fmt.Scanln(&a, &b, &c, &d, &e)
	
	if a < b && b < c && c < d && d < e {
		fmt.Println("Stabil Naik")
	} else if a > b && b > c && c > d && d > e  {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak Stabil") 
	}
}