package main

import "fmt"

func main() {

	cobaSatuDefer()
	fmt.Println()
	cobaBeberapaDefer()
}

func cobaSatuDefer() {
	defer fmt.Println("Halo") // Keyword defer digunakan untuk mengakhirkan statement.
	fmt.Println("Selamat Datang")
	/**
	Pada kode di atas, fmt.Println("Halo") di-defer,
	hasilnya string "Halo" akan muncul setelah "Selamat Datang" .
	*/
}

func cobaBeberapaDefer() {
	defer fmt.Println("Satu")
	defer fmt.Println("Dua")
	defer fmt.Println("Tiga")
	fmt.Println("Empat")
	fmt.Println("Lima")
	// Ketika ada banyak statement yang di-defer,
	// maka statement tersebut akan dieksekusi di akhir secara berurutan.
}
