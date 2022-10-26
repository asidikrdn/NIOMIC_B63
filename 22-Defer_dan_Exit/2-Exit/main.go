package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Teks ini dibuat sebelum exit, tapi menggunakan defer")
	fmt.Println("Teks ini dibuat sebelum exit")
	os.Exit(404) // akan tampil 'exit status 404`
	fmt.Println("Teks ini dibuat setelah exit")

	/**
	Meskipun defer fmt.Println("Teks ini dibuat sebelum exit, tapi menggunakan defer")
	ditempatkan sebelum os.Exit() , statement tersebut tidak akan dieksekusi,
	karena di-tengah fungsi program dihentikan secara paksa.
	*/
}
