package main
import "fmt"

func main() {
	var nilai int
	var tugas bool
	
	fmt.Scanln(&nilai)
	fmt.Scanln(&tugas)
	
	if nilai > 55 && tugas {
		fmt.Println(nilai, true)
	} else {
		fmt.Println(nilai, false)
	}
}