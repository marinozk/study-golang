package main
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	var hasil int = 1	
	
	for i := a; i > 0; i-- {
		hasil = hasil * i
	}
	fmt.Println(hasil)
}