package main
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	var hasil bool
	hasil = a % 2 == 0
	fmt.Println(hasil)
}