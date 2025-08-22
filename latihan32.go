package main
import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	var jam, menit, detik, sisa int
	
	jam = t / 3600
	sisa = t % 3600	
	menit = sisa / 60
	detik = sisa % 60
	fmt.Println(jam, "jam", menit, "menit", detik, "detik")
}