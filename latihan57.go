package main
import "fmt"

func main() {
	var bulan1, bulan2 float64
	fmt.Scan(&bulan1, &bulan2)
	
	if bulan1 < bulan2 {
		bulan2 = bulan2 - bulan1
		fmt.Print("Peningkatan sebesar ", bulan2)
	} else if bulan1 > bulan2 {
		bulan1 = bulan1 - bulan2
		fmt.Print("Penurunan sebesar ", bulan1)
	} else if bulan1 == bulan2 {
		fmt.Print(bulan1)
	}
}