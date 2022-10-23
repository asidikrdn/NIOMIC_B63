package main

import "fmt"

func main() {
	// Variabel biasa
	var a int = 10
	x := "hello world"

	// Variabel pointer
	var b *int = &a
	y := &x

	fmt.Println("Nilai dari variabel a :", a)
	fmt.Println("Alamat dari variabel a :", b)
	fmt.Println("Nilai dari variabel x :", x)
	fmt.Println("Alamat dari variabel x :", y)

	fmt.Println()

	// Nilai dari pointer tidak akan berubah, meskipun nilai variabelnya berubah
	a = 25
	x = "Halo Dunia"
	fmt.Println("Nilai dari variabel a :", a)
	fmt.Println("Alamat dari variabel a :", b)
	fmt.Println("Nilai dari variabel x :", x)
	fmt.Println("Alamat dari variabel x :", y)
}
