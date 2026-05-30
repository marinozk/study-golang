package main
import "fmt"

type tabhimp [100]int

func main() {
	var himpunan tabhimp
	var i int

	isihimp(&himpunan, &i)
	fmt.Println(validasi(himpunan, i))
}

func isihimp(t *tabhimp, i *int) {
	var n int
	fmt.Print("Masukan jumlah data : ")
	fmt.Scan(&n)
	
	for *i = 0; *i < n; *i++ {
		fmt.Scan(&t[*i])
	}
}

func validasi(t tabhimp, i int) string {
	for x := 0; x <= i; x++ {
		for y := x + 1; y <= i; y++ {
			if t[x] == t[y] {
				return "Tidak valid"
			}
		}
	}
	return "Valid"
}