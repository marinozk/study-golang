package main
import "fmt"

func main() {
	var r float64
	const pi = 22.0 / 7.0
	
	fmt.Scanln(&r)
	fmt.Println(4 * pi * r * r)
}