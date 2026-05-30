package main
import "fmt"

const pi = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}
func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}
func display(m1, m2 float64) {
	if m2 == m1 {
		fmt.Println("BALANCE")
	} else {
		fmt.Println(m2 - m1)
	}
}

func main() {
	var r, t1, t2, m1, m2 float64
	fmt.Scanln(&r)
	fmt.Scanln(&t1, &m1)
	fmt.Scanln(&t2, &m2)
	
	massa1 := massa(r, t1, m1)
	massa2 := massa(r, t2, m2)
	display(massa1, massa2)
}