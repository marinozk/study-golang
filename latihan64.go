package main
import "fmt"

func diskon(total float64, member bool) float64 {
	if total > 100000 && member == true {
		return total * 0.9
	} else if total > 100000 && member == false {
		return total * 0.95
	}
	return total
}
func main() {
	var total float64
	var member bool
	fmt.Scan(&total, &member)
	
	hasil := diskon(total, member)
	fmt.Print(hasil)
}
