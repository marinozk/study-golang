package main
import "fmt"

func main() {
	var saldo int
	fmt.Scanln(&saldo)
	var wd int 
	
	for fmt.Scan(&wd); wd != 0; fmt.Scan(&wd) {
		saldo = saldo + wd
	}
	fmt.Println(saldo)
} 