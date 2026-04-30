package main
import "fmt"

func main() {
	var a rune
	fmt.Scanf("%c", &a)
	var hasil bool
	
	hasil = (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')== true
	fmt.Print(hasil)	
}