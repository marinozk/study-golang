package main
import "fmt"

func faktorial(n int,faktor *int) {
	*faktor = 1
	for i := 1; i <= n; i++ {
		*faktor *= i
	}
}

func permutasi(x, y int,p *int) {
	var fx, fxy int
	faktorial(x, &fx)
	faktorial(x-y, &fxy)
	*p = fx / fxy
}
func kombinasi(x,y int, k *int) {
	var fx, fy, fxy int
	faktorial(x, &fx)
	faktorial(y, &fy)
	faktorial(x-y, &fxy)
	*k = fx / (fy * fxy)
}

func main() {
	var a,b,c,d int
	var h1,h2,h3,h4 int
	fmt.Scan(&a,&b,&c,&d)
	
	permutasi(a, c, &h1)
	kombinasi(a, c, &h2)
	permutasi(b, d, &h3)
	kombinasi(b, d, &h4)
	fmt.Println(h1, h2)
	fmt.Println(h3, h4)
}
