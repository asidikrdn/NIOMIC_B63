package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Masukkan angka :")
	fmt.Scanf("%s", &input) // mengambil inputan dari user dan menyimpannya ke variabel input

	var number int
	var err error
	number, err = strconv.Atoi(input) // mengkonversi variabel input ke tipe numerik
	/*
		Fungsi tersebut mengembalikan 2 data, yang kemudian akan ditampung oleh number dan err
		Data pertama ( number ) akan berisi hasil konversi. Dan data kedua err ,
		akan berisi informasi errornya (jika memang terjadi error ketika proses konversi).
	*/

	// Setelah itu dilakukan pengecekkan, ketika tidak ada error, number ditampilkan.
	// Dan jika ada error, input ditampilkan beserta pesan errornya.
	if err == nil {
		fmt.Println(number, "adalah angka")
	} else {
		fmt.Println(input, "bukanlah angka")
		fmt.Println(err.Error()) // Pesan error bisa didapat dari method Error() milik tipe error
	}
}
