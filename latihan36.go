package main
import "fmt"

func main() {
	var huruf string
	fmt.Scanln(&huruf)
	
	if huruf == "b" || huruf == "c" || huruf == "d" || huruf == "f" || huruf == "g" || huruf == "h" || huruf == "j" || huruf == "k" || huruf == "l" || huruf == "m" || huruf == "n" || huruf == "p" || huruf == "q" || huruf == "r" || huruf == "s" || huruf == "t" || huruf == "v" || huruf == "w" || huruf == "x" || huruf == "y" || huruf == "z" ||
		huruf == "B" || huruf == "C" || huruf == "D" || huruf == "F" || huruf == "G" || huruf == "H" || huruf == "J" || huruf == "K" || huruf == "L" || huruf == "M" || huruf == "N" || huruf == "P" || huruf == "Q" || huruf == "R" || huruf == "S" || huruf == "T" || huruf == "V" || huruf == "W" || huruf == "X" || huruf == "Y" || huruf == "Z" {
			fmt.Println("konsonan")
	}	else {
			fmt.Println("Bukan konsonan")
	}

}