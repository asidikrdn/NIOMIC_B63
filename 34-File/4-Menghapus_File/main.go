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

// membuat fungsi menghapus file
func deleteFile(pathFile string) {
	err := os.Remove(pathFile)
	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil dihapus")
}

func main() {
	deleteFile("C:/Users/sidik/Documents/test1.txt")
}
