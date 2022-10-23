package main

import "fmt"

func main() {
	// Panggil fungsi 1
	tampilkanPesan()

	fmt.Println()

	// Panggil fungsi 2
	fmt.Println(pesan())

	fmt.Println()

	// Panggil fungsi 3
	fmt.Println(jumlahkan(5, 7))

}

// Membuat function
func tampilkanPesan() {
	fmt.Println("Hello world")
}

// Membuat function dengan return value
func pesan() string {
	return "Ini Pesan"
}

// Membuat function dengan return value dan parameter
func jumlahkan(x int, y int) int {
	return x + y
}
