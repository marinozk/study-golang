package main
import "fmt"

func main() {
	var total int
	var diskon, kartu, cashback bool
	fmt.Scan(&total, &kartu)
	
	diskon = total >= 100000
	cashback = total >= 200000 && kartu == true
	fmt.Println("kartu?", kartu, "\n", "diskon?", diskon,"\n", "cashback?", cashback)
}