package main
import "fmt"

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m, &x, &y)
	var hasil int
	
	for n >= x && m >= y {
		n = n - x
		m = m - y
		hasil++
		
	}
	fmt.Println(hasil)
}