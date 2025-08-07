package main
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	
	for i := 1; i <= a; i++{
		fmt.Println(i)
	}
}