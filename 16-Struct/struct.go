package main

import "fmt"

func main() {
	// Sekilas struct disini berfungsi layaknya sebuah class pada bahasa pemrograman lain

	// Membuat sebuah variabel bertipe-data struct mahasiswa
	var mhs1 mahasiswa
	mhs1.nama = "Ahmad"
	mhs1.prodi = "Teknik Informatika"
	mhs1.umur = 23

	// Membuat sebuah variabel bertipe-data struct mahasiswa
	var mhs2 mahasiswa
	mhs2.nama = "Sidik"
	mhs2.prodi = "Sistem Informasi"
	mhs2.umur = 22

	fmt.Println("Nama :", mhs1.nama)
	fmt.Println("Umur :", mhs1.umur)
	fmt.Println("Prodi :", mhs1.prodi)
	fmt.Println()
	fmt.Println("Nama :", mhs2.nama)
	fmt.Println("Umur :", mhs2.umur)
	fmt.Println("Prodi :", mhs2.prodi)
}

// Membuat struct
type mahasiswa struct {
	nama  string
	prodi string
	umur  int
}
