package main
import "fmt"

func zoomout(p, l int, skala int) {
	p = p * skala
	l = l * skala
	fmt.Print(p, l)
}
func zoomin(p, l int, skala int) {
	p = p / skala
	l = l / skala
	fmt.Print(p, l)
}

func main() {
	var p, l, skala int
	var s string
	fmt.Scanln(&p, &l)
	fmt.Scanln(&s, &skala)
	if s == "+" {
		zoomout(p, l, skala)
	}else if s == "-" {
		zoomin(p, l, skala)
	}
}