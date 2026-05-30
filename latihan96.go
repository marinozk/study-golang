package main
import "fmt"

const nmax = 37
type thimpunan struct {
	anggota [nmax] int
	panjang int
}

func main() {
	var set1, set2 thimpunan
	
	fmt.Print("Masukan anggota himpunan 1 : ")
	inputhimp(&set1)
	fmt.Print("Masukan anggota himpunan 2 : ")
	inputhimp(&set2)
	fmt.Print("Himpunan 1 = Himpunan 2? ")
	fmt.Print(sama(set1, set2))
}	

func ada(himp thimpunan, n int) bool {
	for i := 0; i < himp.panjang; i++ {
		if n == himp.anggota[i] {
			return false
		}
	}
	return true
}

func inputhimp(himp *thimpunan) {
	var i, n int
	fmt.Scan(&n)
	for ada(*himp, n) {
		himp.anggota[i] = n
		himp.panjang++
		i++
		fmt.Scan(&n)
	}
}

func urut(himp *thimpunan) {
	var temp int

	for i := 0; i < himp.panjang-1; i++ {
		idx := i
		for j := i + 1; j < himp.panjang; j++ {
			if himp.anggota[idx] < himp.anggota[j]	 {
				idx = j
			}
		}
		temp = himp.anggota[idx]
		himp.anggota[idx] = himp.anggota[i]
		himp.anggota[i] = temp
	}
}

func sama(himp1, himp2 thimpunan) bool {
	var i int
	var sama bool
	sama = himp1.panjang == himp2.panjang
	urut(&himp1)
	urut(&himp2)
	for i < himp1.panjang && sama {
		sama = himp1.anggota == himp2.anggota
		i++
	}
	return sama
}