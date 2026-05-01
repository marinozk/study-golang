package main
import "fmt"

func main() {
	var x int
	fmt.Scanf("%c", &x)
	
	if x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z' {	
		fmt.Print("Alphabet")
	} else {
		fmt.Print("Bukan Alphabet")
	}
}