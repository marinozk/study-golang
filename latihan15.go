package main
import "fmt"

func main() {
	var a, b int
	 fmt.Scanln(&a, &b)
	 for x:= a; x <= b; x++ {
		 fmt.Println(x)
	 }
}