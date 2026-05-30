package main
import "fmt"

const bmax = 100
type tabBunga [bmax]string

func main() {
	var daftarBunga tabBunga
	daftarBunga[0] = "Mawar"
	daftarBunga[1] = "Melati"
	daftarBunga[2] = "Anggrek"
	daftarBunga[3] = "Tulip"
	n := 4 

	fmt.Println("Data awal:", daftarBunga[:n])

	posisi := cari(daftarBunga, n, "Melati")
	fmt.Printf("Melati ditemukan di indeks: %d\n\n", posisi)

	fmt.Println("Silakan ketik nama bunga baru untuk menggantikan 'Mawar':")
	rename(&daftarBunga, n, "Mawar") 
	fmt.Println("Data setelah rename:", daftarBunga[:n], "\n")

	fmt.Println("Menghapus bunga 'Anggrek'...")
	deleteB(&daftarBunga, &n, "Anggrek")
	fmt.Println("Data setelah delete:", daftarBunga[:n])
}

func cari(B tabBunga, n int, bunga string) int {
	var i, found int
	found = -1
	i = 0
	
	for i < n && found == -1 {
		if B[i] == bunga {
			found = i
		}
		i += 1
	}
	return found
}

func rename(B *tabBunga, n int, x string) {
	var found int
	found = cari(*B, n, x)
	if found == -1 {
		fmt.Print("Bunga tidak ditemukan")
	} else {
		fmt.Scan(&B[found])
	}
}

func deleteB(B *tabBunga, n *int, x string) {
	var i, found int
	found = cari(*B, *n, x)
	if found == -1 {
		fmt.Print("Bunga tidak ditemukan")
	} else {
		i = found
		for i <= *n-2 {
			B[i] = B[i + 1]
			i += 1
		}
		*n -= 1
	}
}