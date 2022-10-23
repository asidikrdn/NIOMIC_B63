package main

import "fmt"

func main() {
	var rata_rata = hitungRata2(6, 7, 2, 6, 9, 10, 3, 5, 8, 2)
	var pesan = fmt.Sprintf("Rata-ratanya adalah %.2f", rata_rata)
	fmt.Println()
	fmt.Println(pesan)
}

// Membuat fungsi variadic
func hitungRata2(bil ...int) float64 {
	var total int = 0
	// Melakukan perulangan menggunakan for range
	for index, bil := range bil {
		total += bil
		fmt.Printf("bilangan ke-%d adalah %d \n", index, bil)
	}

	var rata2 = float64(total) / float64(len(bil))

	return rata2
}
