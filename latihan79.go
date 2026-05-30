package main
import "fmt"

func pangkat(n int) int {
	if n == 1 {
		return 2
	} else {
		return 2 * pangkat(n-1)
	}
}
func main() {
	var n int 
	fmt.Scan(&n)
	fmt.Println(pangkat(n))
}	