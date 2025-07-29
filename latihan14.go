package main

import "fmt"

func main() {
	var a rune
	fmt.Scanf("%c", &a)

	var hasil bool
	hasil = a >= 'A' && a <= 'Z'

	fmt.Println(hasil)
}
