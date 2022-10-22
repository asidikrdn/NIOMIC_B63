package main

import "fmt"

func main() {
	// Cara 1, dengan menuliskan tipe data
	var namadepan string = "John"

	// Cara 2, tanpa menuliskan tipe data
	var namabelakang = "Doe"

	// Cara 3, tanpa menuliskan var
	namatengah := "bin"

	fmt.Printf("Halo %s %s %s", namadepan, namatengah, namabelakang) // penulisan kode untuk menampilkan nilai variabel mirip seperti penulisan pada bahasa C/C++

	fmt.Println() // Membuat baris baru

	// Mengubah nilai variabeel
	namabelakang = "Syamsudin"

	fmt.Printf("Halo %s %s %s", namadepan, namatengah, namabelakang)
}
