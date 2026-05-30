package main
import "fmt"

func faktorialx(x int) int {
	hasil := 1
	for i := 1; i <= x; i++ {
		hasil *= i
	}
	return hasil
}

func faktorialy(y int) int {
	hasil := 1
	for i := 1; i <= y; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(x, y int) int {
	if x >= y {
		return faktorialx(x) / faktorialy(x-y)
	}
	
	return faktorialx(y) / faktorialy(y-x)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	
	hasil1 := faktorialx(x)
	hasil2 := faktorialy(y)
	hasil3 := permutasi(x, y)
	fmt.Print(hasil1, hasil2, hasil3)
}	