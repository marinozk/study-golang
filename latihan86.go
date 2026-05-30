package main
import "fmt"

type rectangle struct {
	length, width int
	color string 
	property geometry
}

type geometry struct{
	area, perimeter int
}

func isidata(persegi *rectangle) {
	fmt.Scanln(&persegi.length)
	fmt.Scanln(&persegi.width)
	fmt.Scanln(&persegi.color)
}

func hitung(persegi *rectangle) {
	persegi.property.area = persegi.length * persegi.width
	persegi.property.perimeter = (persegi.length * 2) + (persegi.width * 2)
	fmt.Println(persegi.property.area)
	fmt.Println(persegi.property.perimeter)
}

func main() {
	var persegi rectangle
	isidata(&persegi)
	hitung(&persegi)
}