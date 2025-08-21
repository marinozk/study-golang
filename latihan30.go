package main
import "fmt"

func main() {
	var t int 
	fmt.Scanln(&t)
	var total int
	
	for total < t {
		var v int 
		fmt.Scanln(&v)
		total = total + v
		fmt.Println(total >= t)
	} 
}