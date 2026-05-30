package main
import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

func jarak(p1, p2 titik) float64 {
	var jarak float64
	jarak = akar((p1.x - p2.x)*(p1.x - p2.x) + (p1.y - p2.y)*(p1.y - p2.y))
	return jarak
}
func akar(x float64) float64{
	return math.Sqrt(x)
}

func main() {
	var t1, t2 titik
	fmt.Scanln(&t1.x, &t2.x)
	fmt.Scanln(&t1.y, &t2.y)
	fmt.Println(jarak(t1, t2))
}