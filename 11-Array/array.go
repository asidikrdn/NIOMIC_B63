package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "anggur", "melon"} // Membuat array buah yang berisi 4 element string
	fmt.Printf("Jumlah element : %d \n", len(buah))          // Menampilkan maxlength of array buah
	fmt.Printf("Isi element : %s \n", buah)                  // Menampilkan isi dari array buah

	fmt.Println()

	// Apabila isinya kita kurangi, panjang array tetap akan sesuai jumlah element maksimalnya
	var buah2 = [4]string{"apel", "jeruk", "anggur"} // Membuat array buah yang hanya berisi 3 element string, namun dengan max elementnya tetap 4
	fmt.Printf("Jumlah element : %d \n", len(buah2)) // Hasilnya akan tetap 4 walau array buah hanya berisi 3 element
	fmt.Printf("Isi element : %s \n", buah2)

	fmt.Println()

	// Merubah element array
	var buah3 = [4]string{"apel", "jeruk", "anggur", "melon"}
	fmt.Printf("Jumlah element : %d \n", len(buah3))
	fmt.Printf("Isi element : %s \n", buah3)
	buah3[2] = "mangga" // merubah element index ke-2
	fmt.Printf("Jumlah element : %d \n", len(buah3))
	fmt.Printf("Isi element : %s", buah3)
}
