package main
import "fmt"

const nmhs = 2022
type mahasiswa struct {
	sid int
	nama string
	gpa float64
}
type arrmhs [nmhs]mahasiswa

func main() {
	var data arrmhs
	data[0] = mahasiswa {113210689, "Harith", 1.56}
	data[1] = mahasiswa {113212624, "Johnson", 3.19}
	data[2] = mahasiswa {113211834, "Kimmy", 1.32}
	data[3] = mahasiswa {113212925, "Chou", 3.68}
	data[4] = mahasiswa {113210520, "Grock", 1.45}
	data[5] = mahasiswa {113210223, "Lunox", 1.89}
	data[6] = mahasiswa {113212819, "Karrie", 1.05}
	data[7] = mahasiswa {113211273, "aldous", 2.46}
	data[8] = mahasiswa {113211643, "Franco", 1.60}
	data[9] = mahasiswa {113211992, "Selena", 3.50}
	n := 10
	
	fmt.Println("Data Mahasiswa : ")
	for i := 0; i < n; i++ {
		fmt.Println(data[i],)
	}
	mengurutkan(&data, &n)
	fmt.Println(" ")
	fmt.Println("Data Ascending : ")
	for i := 0; i < n; i++ {
		fmt.Println(data[i])
	}
}

func mengurutkan(t *arrmhs, n *int) {
	var idx, pass, i int 
	var temp mahasiswa
	pass = 1
	
	for pass < *n {
		idx = pass-1
		i = pass
		for i < *n {
			if t[idx].gpa > t[i].gpa	 {
				idx = i
			}
			i += 1
		}
		temp = (*t)[pass-1]
		(*t)[pass-1] = (*t)[idx]
		(*t)[idx] = temp
		pass += 1
	}
}

