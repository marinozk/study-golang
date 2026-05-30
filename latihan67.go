package main
import "fmt"

func hitungluas(a, t int, luas *float64) {
	*luas = float64(a * t) / 2
}

func main() {
	var n, t, a int
	var luas float64
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &t)
		hitungluas(a, t, &luas)
		fmt.Println(luas)
	}
}