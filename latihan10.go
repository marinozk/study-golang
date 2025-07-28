package main
import "fmt"

func main() {
	var a, b, c float32
	
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a + (a * 5/100))
	fmt.Println(b + (b * 5/100))
	fmt.Println(c + (c * 5/100))
}