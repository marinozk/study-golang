package main
import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(x,y int) int {
	return faktorial(x) / faktorial(x-y)
}

func kombinasi(x, y int) int {
	return faktorial(x) / (faktorial(y) * faktorial(x-y))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}