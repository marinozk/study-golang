package main
import "fmt"

func mengurutkan(a, b int, x, y *int){
	if a >= b {
	*x = b
	*y = a
	}else {
	*x = a
	*y = b
	}	
}

func main() {
	var a, b int
	var x, y int
	fmt.Scan(&a, &b)
	
	mengurutkan(a, b, &x, &y)
	fmt.Print(x, y)
}