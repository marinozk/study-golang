package main
import "fmt"

func main() {
	var x, y float32
	fmt.Scanln(&x, &y)
	
	if x < y {
		fmt.Println("Peningkatan Sebesar", y - x)
	}
	if x > y {
		fmt.Println("Penurunan Sebesar", x - y)
	}
	if x == y {
		fmt.Println("Tetap")
	}
}