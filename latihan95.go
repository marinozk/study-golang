package main
import "fmt"

type tabint [100]int

func main() {
	var data tabint
	var n int
	
	input(&data, &n)
	insertionsort(&data, n)
	fmt.Print(median(data, n))
}

func input(t *tabint, i *int) {
	var n int
	fmt.Scan(&n)
	for *i = 0; *i < n; *i++ {
		fmt.Scan(&t[*i])
	}
}

func insertionsort(t *tabint, n int) {
	var pass, i, temp int
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = (*t)[pass]
		for i > 0 && temp < (*t)[i-1] {
			(*t)[i] = (*t)[i-1]
			i--
		}
		(*t)[i] = temp
		pass++
	}
}

func median(t tabint, n int) float64 {
	jum := 0
	for i := 0; i < n; i++ {
		jum += t[i]
	}
	return float64(jum) / float64(n)
}
