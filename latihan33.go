package main
import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	
	var hasil bool = true
	var dsebelum int = a % 10
	a = a / 10
	
	for a > 0 {
		var dsekarang int = a % 10
		hasil = hasil && ((dsekarang - dsebelum == 1) || (dsebelum - dsekarang == 1))
		dsebelum = dsekarang
		a = a/10
	}
	fmt.Println(hasil)
}