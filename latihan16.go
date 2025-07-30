package main
import "fmt"

func main() {
	var a,d,total int
	
	fmt.Scanln(&a)
	for i:= 0; i <a; i++ {
		fmt.Scanln(&d)
		
		var DigitAkhir int = d % 10
		var DigitAwal int = d
		for DigitAwal >= 10 {
			DigitAwal /= 10
		} 
		total += DigitAkhir + DigitAwal
	}
	fmt.Println(total)
}