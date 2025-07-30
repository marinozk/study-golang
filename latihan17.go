package main
import "fmt"

func main() {
	var d int
	var s string
	
	fmt.Scanln(&d)
	fmt.Scanln(&s)
	
	for i := 0; i < d; i++ {
		fmt.Println(s)
	}
}