package main
import "fmt"

type balok struct {
	p, l, t int
	volume, luas int 
}

func luas(x *balok) {
	x.luas = (x.p * x.l + x.p * x.t + x.t * x.l)
}

func volume(x *balok) {
	x.volume = (x.p * x.t * x.l)
}

func main() {
	var block balok
	fmt.Scan(&block.p, &block.l, &block.t)
	
	luas(&block)
	volume(&block)
	
	fmt.Println(block.luas)
	fmt.Println(block.volume)
}