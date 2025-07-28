package main
import "fmt"

func main() {
	var panjang, lebar int
	
	fmt.Scanln(&panjang, &lebar)
	fmt.Println((panjang+panjang) + (lebar+lebar))
	fmt.Println(panjang * lebar)
}