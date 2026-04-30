package main
import "fmt"

func main() {
	var uang int
	var card bool
	var diskon int = 100000
	var cb int = 200000
	
	fmt.Scanln(&uang, &card)
	if uang >= diskon {
		uang = uang - (uang * 10/100)
		adadiskon := true
	}
	if card == true && uang >= cb {
		hcb = hdis + 75000
		cashback := true
	}
	fmt.Printf("Kartu? %v\n", card)
	fmt.Printf("Diskon? %v\n", adadiskon)
	fmt.Printf("Cashback? %v\n", cashback)
	fmt.Printf("Rp. %d\n", )
	
}