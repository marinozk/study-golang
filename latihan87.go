package main
import "fmt"

type palindrom [256] int

func pengisian(bil *palindrom, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&bil[i])
	}
}

func reverse(bil *palindrom, n int) {
	var i, j int
	for i,j = 0, n-1; i < j; i, j = i+1, j-1 {
		bil[i], bil[j] = bil[j], bil[i]
	}
}

func isipalindrom(bil palindrom, n int) bool {
	var i,j int
	for i,j = 0, n-1; i < j; i, j = i+1, j-1 {
		if bil[i] != bil[j] {
		return false
		}
	}
	return true
}

func main() {
	var bil palindrom
	var n int
	
	fmt.Scan(&n)
	pengisian(&bil, n)
	fmt.Println(bil [:n])
	
	if isipalindrom(bil, n) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan Palindrom")
	}

	reverse(&bil, n)
	fmt.Println(bil[:n])
}

