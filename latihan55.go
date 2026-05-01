package main
import "fmt"

func main() {
	var x string
	fmt.Scan(&x)
	
if x == "a" || x == "i" || x == "u" || x == "e" || x == "o" || x == "A" || x == "I" || x == "U" || x == "E" || x == "O" {
		fmt.Println("Bukan Konsonan")
	} else {
		fmt.Println("Konsonan")
	}
}