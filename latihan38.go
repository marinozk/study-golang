package main
import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	
	if a == b && b == a {
		fmt.Println("Segitiga sama sisi")
	} else if a == c && b != a {
		fmt.Println("Segitiga sama kaki")
	} else if a != b && b != c {
		fmt.Println("Segitiga semabarang")
	}
}