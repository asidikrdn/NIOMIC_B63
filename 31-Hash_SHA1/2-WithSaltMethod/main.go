package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	text := "Ini rahasia"
	fmt.Printf("text asli : %s \n\n", text)

	// Melakukan hasing pertama
	hashed1, salt1 := hasingMenggunakanSalt(text)
	fmt.Printf("Hashed 1 : %s\n\n", hashed1)
	_ = salt1
	// Melakukan hasing kedua
	hashed2, salt2 := hasingMenggunakanSalt(text)
	fmt.Printf("Hashed 2 : %s\n\n", hashed2)
	_ = salt2
	// Melakukan hasing ketiga
	hashed3, salt3 := hasingMenggunakanSalt(text)
	fmt.Printf("Hashed 3 : %s\n\n", hashed3)
	_ = salt3
	// Meskipun text aslinya sama, namun nilai hashingnya akan berbeda-beda karena terdapat proses salting berbasis waktu.

	fmt.Println()

	// Melakukan hasing pertama
	hashed1, salt1 = hasingMenggunakanSaltStatic(text)
	fmt.Printf("Hashed 1 : %s\n\n", hashed1)
	_ = salt1
	// Melakukan hasing kedua
	hashed2, salt2 = hasingMenggunakanSaltStatic(text)
	fmt.Printf("Hashed 2 : %s\n\n", hashed2)
	_ = salt2
	// Melakukan hasing ketiga
	hashed3, salt3 = hasingMenggunakanSaltStatic(text)
	fmt.Printf("Hashed 3 : %s\n\n", hashed3)
	_ = salt3
	// Jika saltnya menggunakan static, maka mau dihasing berapa kali pun akan sama hasilnya (karena text asli dan saltnya selalu sama)
}

// membuat fungsi hashing dengan salt
func hasingMenggunakanSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("text: '%s', salt: '%s'", text, salt)
	fmt.Println(saltedText)

	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func hasingMenggunakanSaltStatic(text string) (string, string) {
	salt := "12345"
	saltedText := fmt.Sprintf("text: '%s', salt: '%s'", text, salt)
	fmt.Println(saltedText)

	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
