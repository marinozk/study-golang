package main
import ("fmt")

func main() {
	var a int
	var b int
	var c int
	var hasil int 
	
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	
	hasil = b * c + a
	fmt.Println(hasil)
}
