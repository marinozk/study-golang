package main
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	var hasil bool
	
	for i := 1; i <= a; i++ {
		hasil = a % i == 0
		fmt.Println(i, hasil)
	}
}