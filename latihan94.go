package main
import "fmt"

const maks = 30
type intarr struct {
	tabint [maks]int
	n int 
}
var array1, array2 intarr

func main() {
	fmt.Println("Masukan array pertama")
	inputarray(&array1)
	fmt.Println(" ")
	fmt.Println("Masukan array kedua")
	inputarray(&array2)
	
	fmt.Println(" ")
	sortarray(&array1)
	sortarray(&array2)
	fmt.Println("Hasil sort array pertama : ", array1.tabint[:array1.n])
	fmt.Println("Hasil sort array pertama : ", array2.tabint[:array2.n])
	
	fmt.Println("Similar : ", similar(array1, array2))
}

func inputarray(t *intarr) {
	fmt.Print("Jumlah Elemen : ")
	fmt.Scan(&t.n)
	for i := 0; i < t.n; i++ {
		fmt.Print("Elemen ke ", i, ": ")
		fmt.Scan(&t.tabint[i])
	}		
}

func sortarray(t *intarr) {
	var idx, i, pass, temp int
	
	pass = 1
	for pass < t.n {
		idx = pass-1
		i = pass
		for i < t.n {
		if t.tabint[idx] > t.tabint[i] {
			idx = i
			}
			i++
		}
		temp = t.tabint[pass-1]
		t.tabint[pass-1] = t.tabint[idx]
		t.tabint[idx] = temp
		pass++
	}
}

func similar(t1, t2 intarr) bool {
	
	if t1.n != t2.n {
		return false
	}
	
	for i := 0; i < t1.n; i++ {
		if t1.tabint[i] != t2.tabint[i] {
			return false
		}
	}
	return true
}