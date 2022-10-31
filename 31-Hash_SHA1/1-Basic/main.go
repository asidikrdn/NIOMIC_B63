package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	text := "Ini rahasia"

	// membuat variabel sha yang merupakan turunan dari struct sha1, yakni objek bertipe hash.Hash, memiliki dua buah method Write() dan Sum().
	sha := sha1.New()

	// Menge-set data yang akan di-hash. Data harus dalam bentuk []byte.
	sha.Write([]byte(text))

	// Mengeksekusi proses hash, menghasilkan data yang sudah di-hash dalam bentuk []byte. Method ini membutuhkan sebuah parameter, isi dengan nil.
	encrypted := sha.Sum(nil)

	// menyimpan bentuk hexadesimal dari data yang sudah di hash ke dalam sebuah variabel
	encryptedStr := fmt.Sprintf("%x", encrypted)

	// menampilkan hasil hashing dengan bentuk asli
	fmt.Println(encrypted)
	// menampilkan hasil hashing dengan bentuk hexadesimal
	fmt.Println(encryptedStr)
	// menampilkan hasil hashing dengan bentuk hexadesimal
	fmt.Printf("%x \n", encrypted)
}
