package main
import "fmt"

func main() {
	var d, d1, d2 int
	
	fmt.Scanln(&d)
	d1 = d / 10
	d2 = d % 10
	fmt.Println(d1,d1,d2,d2)
	
	
}