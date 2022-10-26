package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var nama string
	fmt.Print("Masukkan nama anda : ")
	fmt.Scanf("%s", &nama)

	// Mengambil nilai kembalian dari fungsi validasi dan menyimpannya ke dalam variabel valid dan err
	if valid, err := validasi(nama); valid { // apabila valid bernilai true, maka jalankan
		fmt.Println("Halo", nama)
	} else {
		fmt.Println(err.Error())
	}

	/**
	Jika kode diatas sulit dipahami, bisa dilihat kode berikut yang fungsinya sama

	valid, err := validasi(nama)

	if valid {
		fmt.Println("Halo", nama)
	} else {
		fmt.Println(err.Error())
	}
	*/
}

// membuat fungsi validasi dengan 2 nilai kembalian
func validasi(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input tidak boleh kosong")
	} else {
		return true, nil
	}
}
