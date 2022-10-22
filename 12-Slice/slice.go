package main

import "fmt"

func main() {
	var buah = []string{"apel", "mangga", "jeruk", "anggur"} // Membuat slice buah
	fmt.Printf("Jumlah element : %d \n", len(buah))
	fmt.Printf("Isi element : %s \n", buah)

	fmt.Println()

	buah = append(buah, "melon") // Menambahkan element baru pada slice buah
	fmt.Printf("Jumlah element : %d \n", len(buah))
	fmt.Printf("Isi element : %s \n", buah)
}
