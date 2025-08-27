package main
import "fmt"

func main() {
	var digit int
	fmt.Scanln(&digit)

	if digit % 5 == 0 {
		fmt.Println("Kelipatan 5")
	}
	if digit % 3 == 0 {
		fmt.Println("Kelipatan 3")
	}
	if digit % 5 != 0 && digit % 3 != 0 {
		fmt.Print(" ")
	}
}