package main

import (
	"fmt"
	"io"
	"os"
)

// fungsi menampilkan pesan error
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}

// membuat fungsi untuk membaca file
func readFile(pathFile string) {
	// membuka file dengan level akses READ ONLY
	file, err := os.OpenFile(pathFile, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// membaca isi file
	text := make([]byte, 1024) // membuat variabel penampung hasil pembacaan isi file
	for {
		n, err := file.Read(text)

		// jika terdapat error selain `io.EOF`, maka tampilkan pesan error dan hentikan perulangan
		if err != io.EOF {
			if isError(err) {
				break
			}
		}

		// jika isi file sudah habis (0), hentikan perulangan
		if n == 0 {
			break
		}
	}

	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil dibaca, berikut adalah isinya : \n ")
	fmt.Println(string(text))
}

func main() {
	readFile("C:/Users/sidik/Documents/test1.txt")
}
