package main

import "fmt"

func main() {
	// Membuat sebuah map dengan nama harga_makanan
	var harga_makanan = map[string]int{"ayam_goreng": 20000, "sate_ayam": 15000}

	// Menampilkan nilai sebuah element map
	fmt.Printf("Harga ayam goreng : Rp %d\n", harga_makanan["ayam_goreng"])
	fmt.Printf("Harga sate ayam : Rp %d\n", harga_makanan["sate_ayam"])
}
