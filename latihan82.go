package main
import "fmt"

func faktor(bil, n int) {
	if bil > n{
		return
	}
	if n % bil == 0 {
		fmt.Print(bil, " ")
	}
	faktor(bil + 1, n)
}
func main() {
	var n int
	fmt.Scan(&n)
	faktor(1, n)
}