package main
import "fmt"

func hasilnya(x, y, cx, cy int) int {
	x = x - cx
	x = x*x
	y = y - cy
	y = y*y
	hasil := x + y
	return hasil
}

func cek(x, y, cx, cy, cx2, cy2, r, r2 int) string {
	lingkaran1 := hasilnya(x, y, cx, cy) <= r*r
	lingkaran2 := hasilnya(x, y, cx2, cy2) <= r2*r2
	
	if lingkaran1 && lingkaran2{
		return "titik di dalam lingkaran 1 dan 2"
	} else if lingkaran1 && !lingkaran2 {
		return "titik di luar lingkaran 1"
	} else if !lingkaran1 && lingkaran2 {
		return"titik di dalam lingkaran 2"
	} 
	return "titik di luar lingkaran 1 dan 2"	
}

func main() {
	var cx, cy, cx2, cy2, r, r2, x, y int
	fmt.Scanln(&cx, &cy, &r, &cx2, &cy2, &r2, &x, &y)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)
	
	keluaran := cek(x, y, cx, cy, cx2, cy2, r, r2)
	fmt.Println(keluaran)
}