package main

import (
	"fmt"
	"os"
)

// fungsi menampilkan pesan error
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}

// fungsi untuk membuat file
func createFile(path string) {
	// memeriksa apakah file yang ingin dibuat sudah ada
	_, err := os.Stat(path)

	// jika file belum ada, maka buat file baru
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if isError(err) {
			return
		}
		// menutup file setelah dibuat
		defer file.Close()

		// menampilkan teks apabila sebuah file berhasil dibuat
		fmt.Println("===> file berhasil dibuat :", path)
	} else {
		// menampilkan teks apabila sebuah file sudah ada
		fmt.Println("===> file sudah dibuat !")
	}

}

func main() {
	createFile("34-File/test1.txt")
}
