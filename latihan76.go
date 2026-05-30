package main
import (
	"fmt"
	"math"
)

func jarak(a,b,c,d int) float64{
	jx := float64(a - c)
	jy := float64(b - d)
	return jx*jx + jy*jy
}

func didalam(d1, d2, r1, r2 float64) string{
	if d1 <= r1 && d2 <= r2 {
		return "titik di dalam lingkaran 1 dan 2"
	} else if d1 <= r1 && d2 >= r2 {
		return "titik di dalam lingkarang 1"
	} else if d1 >= r1 && d2 <= r2 {
		return "titik di dalam lingkarang 2"
	}
	return "titidak di luar lingkarang 1 dan 2"
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	var cx, cy, cx2, cy2, x, y int 
	var r, r2 float64
	fmt.Scanln(&cx, &cy, &r)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)
	
	d1 := jarak(cx, cy, x, y)
	d2 := jarak(cx2, cy2, x, y)
	a1 := akar (d1)
	a2:= akar (d2)
	
	fmt.Println(didalam(a1, a2, r, r2))
}

