package main
import "fmt"

type tabhimp [100]int

func main() {
	var himp, himp2, hasil tabhimp
	var i, i2, ihasil int

	fmt.Print("Masukan data pertama : ")
	isidata(&himp, &i)
	fmt.Print("Masukan data kedua : ")
	isidata(&himp2, &i2)
	
	gabungan(himp, i, himp2, i2, &hasil, &ihasil)
	fmt.Print("Gabungan { ")
	for i := 0; i < ihasil; i++ {
		fmt.Print(hasil[i])
		if i != ihasil - 1 {
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

func ada(t tabhimp, n int, x int) bool {
	for i := 0; i < n; i++ {
		if t[i] == x {
			return true
		}
	}
	return false
}

func gabungan(t tabhimp, i int,
	t2 tabhimp, i2 int,
	hasil *tabhimp, ihasil *int) {
	*ihasil = 0
	
	for x := 0; x < i; x++ {
		hasil[*ihasil] = t[x]
		(*ihasil)++
	}
	for y := 0; y < i2; y++ {
		if !ada(*hasil, *ihasil, t2[y]) {
			hasil[*ihasil] = t2[y]
			(*ihasil)++
		}
	}
}