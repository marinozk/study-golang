package main
import "fmt"

func segitiga(s1, s2, s3 int) bool {
	return s1 + s2 > s3 && s2 + s3 > s1 && s3 + s1 > s2 == true
}

func main() {
	var s1, s2, s3 int
	fmt.Scan(&s1, &s2, &s3)
	
	hasil := segitiga(s1, s2, s3)
	if hasil == true {
		fmt.Print("segitiga")
	} else {
		fmt.Print("bukan seigitiga")
	}	
}

