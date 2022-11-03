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

// membuat fungsi untuk mengedit file
func writeToFile(pathFile string, textSlice []string) {
	// membuka file dengan level akses READ & WRITE
	file, err := os.OpenFile(pathFile, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	for _, text := range textSlice {
		_, err = file.WriteString(text + "\n")
		if isError(err) {
			return
		}
	}

	// menyimpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil diisi")
}

func main() {
	writeToFile("34-File/test1.txt", []string{"Ahmad Sidik Rudini", "M Rizki Syabani"})
}
