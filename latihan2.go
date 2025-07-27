package main
import ("fmt")

func main () {
	var a, hasil1, mod1, hasil2, mod2, hasil3 int
	
	fmt.Scanln(&a)
	
	hasil1 = a / 10000
	mod1 = a % 10000 
	hasil2 = mod1 / 5000
	mod2 = mod1 % 5000
	hasil3 = mod2 / 1000
	
	fmt.Println(hasil1, "lembar")
	fmt.Println(hasil2, "lembar")
	fmt.Println(hasil3, "lembar")
	
}