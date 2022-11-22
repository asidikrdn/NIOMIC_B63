package library

import "fmt"

type Mahasiswa struct {
	nama string
	umur int
}

func (mhs Mahasiswa) Tampilkan_mahasiswa() {
	mhs.nama = "Sidik"
	mhs.umur = 23
	fmt.Printf("Nama : %s \nUmur : %d \n", mhs.nama, mhs.umur)
}
