package main
import "fmt"

func ganjil(now, n int) {
	if now > n {
	return
	}
	if now % 2 != 0 {
		fmt.Print(now, " ")
	}
	ganjil(now + 1, n)
}
func main() {
	var n int
	fmt.Scan(&n)
	ganjil(1, n)
}