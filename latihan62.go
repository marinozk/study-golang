package main
import "fmt"

func kjam(jam int) int {
	return jam * 3600
} 

func kmenit(menit int) int {
	return menit *60
}

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	
	hasil := kjam(jam) + kmenit(menit) + detik
	fmt.Print("hasil konversi = ", hasil, " detik")
	
}