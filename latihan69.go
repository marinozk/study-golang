package main
import "fmt"

func pecahan(uang int, p10, p2, p1 *int) {
	var sisa int
	*p10 = uang / 10000
	sisa = uang % 10000
	
	*p2 = sisa / 2000
	sisa = sisa % 2000
	
	*p1 = sisa / 1000 
	}
	
func main() {
	var	uang, p10, p2, p1 int
	fmt.Scan(&uang)
	
	pecahan(uang, &p10, &p2, &p1)
	fmt.Println(p10, "lembar 10000")
	fmt.Println(p2, "lembar 2000")
	fmt.Println(p1	, "lembar 1000")
	
}