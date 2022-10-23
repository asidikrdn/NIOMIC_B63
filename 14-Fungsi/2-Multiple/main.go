package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	var x1, x2 = kalibagi(a, b)

	fmt.Printf("Hasil dari %d x %d adalah %d\n", a, b, x1)
	fmt.Printf("Hasil dari %d : %d adalah %d\n", a, b, x2)
}

// Membuat fungsi dengan 2 return value
func kalibagi(x int, y int) (int, int) {
	var hasilKali = x * y
	var hasilBagi = x / y

	return hasilKali, hasilBagi
}
