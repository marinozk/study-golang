package main
import "fmt"

func main() {
		var n, mobil int
		fmt.Scan(&n)
		mobil = 1
		
		if n > 7 {	
			mobil = mobil + (n / 7)
		}
		fmt.Print(mobil)
}