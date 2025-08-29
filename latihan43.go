package main
import "fmt"

func main() {
	var jam1, menit1, jam2, menit2, hasil1, hasil2, durasi, jam, menit int
	fmt.Scanln(&jam1, &menit1, &jam2, &menit2)
	
	hasil1 = jam1 * 60 + menit1 
	hasil2 = jam2 * 60 + menit2 
	
	if hasil2 < hasil1 {
		hasil2 = hasil2 + 12 * 60
	}
	
	durasi = hasil2 - hasil1
	jam = durasi / 60
	menit = durasi % 60
	
	fmt.Println("jam", jam, "menit", menit)
}