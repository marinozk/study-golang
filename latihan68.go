package main
import "fmt"

func nilaiekstrim(n1, n2 int, kecil, besar	 *int) {
	if n1 >= n2 {
	*besar = n1
	*kecil = n2
	} else if n2 >= n1{
	*besar = n2	
	*kecil = n1
	}	
}

func main() {
	var a,b,c,d,kecil,besar int
	var kecil1, kecil2, besar1, besar2, temp int
	fmt.Scan(&a,&b,&c,&d)
	
	nilaiekstrim(a, b, &kecil1, &besar1)
	nilaiekstrim(c, d, &kecil2, &besar2)
	nilaiekstrim(besar1, besar2, &temp, &besar)
	nilaiekstrim(kecil1, kecil2, &kecil, &temp)
	fmt.Print(besar, kecil)
}