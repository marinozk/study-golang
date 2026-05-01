package main
import "fmt"

func main() {
	var a, b, c int
	var segitiga bool
	fmt.Scanln(&a, &b, &c)
	
	segitiga = (a + b > c && b + c > a && a + c > b) && (a > 0) && (b > 0) && (c > 0)
	fmt.Print(segitiga)
}