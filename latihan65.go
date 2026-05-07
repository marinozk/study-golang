package main
import "fmt"

func fibonacci(n int) int{
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	
	a := 0
	b := 1
	total := a + b
	
	for i := 3; i <= n; i++ {
	sukubaru := a + b
	total += sukubaru
	
	a = b
	b = sukubaru
	}
	return total
}

func main() {
	var n int
	fmt.Scan(&n)
	
	hasil := fibonacci(n)
	fmt.Println(hasil)
	
}