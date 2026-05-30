package main
import "fmt"

type teman [100]string

func tulis(t *teman, n int) {
	for i := 0; i<n; i++ {
		fmt.Scanln(&t[i])
	}
}

func cetak(t teman, n int) {
	for i := 0; i<n; i++ {
		fmt.Println(t[i])
	}
}

func main() {
	var n int
	var t teman
	fmt.Scanln(&n)
	tulis(&t, n)
	cetak(t, n)
}