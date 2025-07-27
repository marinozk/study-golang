package main
import "fmt"

func main() {
	var c float32
	
	fmt.Scanln(&c)
	
	fmt.Println(c * 0.8)
	fmt.Println(c * 1.8 + 32)
	fmt.Println(c + 273.15)
	
}