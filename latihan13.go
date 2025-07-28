package main
import "fmt"

func main() {
	var Eo, c, E1 int
	
	fmt.Scanln(&Eo, &c, &E1)
	fmt.Println((Eo-E1) / c)
}