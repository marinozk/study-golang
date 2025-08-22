package main
import "fmt"

func main() {
	var n, c int 
	fmt.Scanln(&n)
	var a, b int = 0, 1
	fmt.Print(a,b )
	
	for i := 2; i <= n; i++{
		c = a + b
		fmt.Print(" ",c)
		a, b = b, c
	} 
}