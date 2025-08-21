package main
import "fmt"

func main() {
	var x int
	fmt.Scanln(&x)
	var total,simpan int
	
	for x > 0 {
		simpan = x % 10
		total = total + simpan
		x = x / 10
	}
	fmt.Println(total)
}