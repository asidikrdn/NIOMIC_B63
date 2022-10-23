package main

import "fmt"

func main() {
	// Mendeklarasikan beberapa variabel sekaliguns yang menginstance ke struct mahasiswa
	var mhs1, mhs2 mahasiswa
	mhs1.nama = "Ahmad"
	mhs1.prodi = "Teknik Informatika"
	mhs1.umur = 23
	mhs2.nama = "Sidik"
	mhs2.prodi = "Sistem Informasi"
	mhs2.umur = 22

	// Mendeklarasikan variabel instance dari struct mahasiswa dan langsung memberikan nilai saat proses deklarasi
	var mhs3 = mahasiswa{"Rudini", "Teknik Elektro", 21}

	mhs1.tampilkanNama()
	mhs1.tampilkanUmur()

	fmt.Println()

	mhs2.tampilkanNama()
	mhs2.tampilkanUmur()

	fmt.Println()

	mhs3.tampilkanNama()
	mhs3.tampilkanUmur()
}

// Membuat struct
type mahasiswa struct {
	nama  string
	prodi string
	umur  int
}

// Membuat method
func (var_object mahasiswa) tampilkanNama() {
	fmt.Println("Nama saya adalah", var_object.nama)
}

func (var_object mahasiswa) tampilkanUmur() {
	fmt.Println("Umur saya adalah", var_object.umur)
}
