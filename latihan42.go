package main
import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)
	
	var maks int = a
	if b > maks {
		maks = b 
	}
	if c > maks {
		maks = c
	}
	if d > maks {
		maks = d
	}
	
	var minm int = a
	if b < minm {
		minm = b
	}
	if c < minm {
		minm = c
	}
	if d < minm {
		minm = d
	}
	fmt.Println(maks, minm	)
}