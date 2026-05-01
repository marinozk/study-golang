package main
import "fmt"

func main() {
	var total, diskon float64
	var clo bool
	fmt.Scanln(&total, &clo)
	
	if clo == true {
		diskon = total * 0.35
		total = total - diskon 
	} 
	fmt.Print(total)
}