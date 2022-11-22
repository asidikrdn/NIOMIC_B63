package main

import "fmt"

var mahasiswa = map[string]int{"Aldo": 182, "Yosep": 178}

func tampil_mahasiswa(mhs map[string]int) {
	tampil := func() {
		for nama, tinggi := range mhs {
			fmt.Println(nama, ":", tinggi, "cm")
		}
	}

	tampil()
}

func main() {
	tampil_mahasiswa(mahasiswa)
}
