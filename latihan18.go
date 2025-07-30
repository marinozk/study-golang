package main
import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var sisi int
		fmt.Scanln(&sisi)

		var luas int = sisi * sisi
		var keliling int = 4 * sisi

		fmt.Println(luas, keliling)
	}
}
