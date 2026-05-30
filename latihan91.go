package main
import "fmt"

type tabhimp [100]int

func main() {
	var himpunan1, himpunan2, hasil tabhimp
	var i1, i2, nhasil int
	
	fmt.Print("Masukan data pertama : ")
	isidata(&himpunan1, &i1)
	fmt.Print("Masukan data kedua : ")
	isidata(&himpunan2, &i2)
	
	irisan(himpunan1, i1, himpunan2, i2, &hasil, &nhasil)
	
	fmt.Print("Irisan { ")
	for i := 0; i < nhasil; i++ {
		fmt.Print(hasil[i])
		if i != nhasil-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print(" }")
}

func isidata(t *tabhimp, i *int) {
	var n int
	fmt.Scan(&n)
	
	for *i = 0; *i < n; *i++ {
		fmt.Scan(&t[*i])
	}
}

func irisan(t tabhimp, i int, t2 tabhimp, i2 int, hasil *tabhimp, nhasil *int) {
	*nhasil = 0
	
	for x := 0; x < i; x++ {
		for y := 0; y < i2; y++ {
			if t[x] == t[y] {
				hasil[*nhasil] = t[x]
				*nhasil++
			}
		}
	}
} 