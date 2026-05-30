package main
import "fmt"

const nmax = 1000
type lulusan struct {
	nama, nim string
	semester int
	eprt, ipk float64
}

type tablulus [nmax]lulusan

func main() {
	var data tablulus
	var n int
	var eprtertinggi, ipkterendah, ratasmt float64
	
	isidata(&data, &n)
	eprtertinggi = maxeprt(data, n)
	ipkterendah = minipk(data, n)
	ratasmt = ratasemester(data, n)
	
	fmt.Printf("EPRT Tertinggi: %.2f\n", eprtertinggi)
	fmt.Printf("IPK Terendah  : %.2f\n", ipkterendah)
	fmt.Printf("Rata Semester : %.2f\n", ratasmt)
}

func isidata(T *tablulus, n *int) {
	var nim string
	*n = 0
	fmt.Scan(&nim)
	for nim != "none" && *n <= nmax {
		T[*n].nim = nim
		fmt.Scan(&T[*n].nama, &T[*n].semester, &T[*n].eprt, &T[*n].ipk)
		*n += 1
		fmt.Scan(&nim)
	}
}

func maxeprt(T tablulus, n int) float64 {
	var maks float64
	var i int
	maks = T[0].eprt
	i = 1
	
	for i < n {
		if maks < T[i].eprt {
			maks = T[i].eprt
		}
		i += 1
	}
	return maks
}

func minipk(T tablulus, n int) float64 {
	var mins lulusan
	var i int
	mins = T[0]
	i = 1
	
	for i < n {
		if mins.ipk > T[i].ipk {
			mins = T[i]
		}
		i += 1
	}
	return mins.ipk
}

func ratasemester(T tablulus, n int) float64 {
	var jum float64
	var i int
	jum = 0
	i = 0
	
	for i < n {
		jum += float64(T[i].semester)
		i += 1
	}
	return jum / float64(n)
}