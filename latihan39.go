package main	
import "fmt"

func main() {
	var angka, hasil int
	fmt.Scanln(&angka)
	
	if angka < 0 {
		hasil = angka * (-1)
		fmt.Println(hasil)
	}
	if angka >= 0 {
		hasil = angka
		fmt.Println(hasil)
	}
}