package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	nama := "Ahmad Sidik Rudini"

	// EncodingToString
	// varHasilEncode := base64.URLEncoding.EncodeToString([]byte(stringYangInginDiEncode))
	// Note : []byte() digunakan untuk meng-casting string yang ingin diencode
	encodedStr := base64.StdEncoding.EncodeToString([]byte(nama))

	sha := sha1.New()

	sha.Write([]byte(encodedStr))

	encrypted := sha.Sum(nil)

	fmt.Println("Text Asli :", nama)
	fmt.Println("Hasil encoding dengan base64 :", encodedStr)
	fmt.Printf("Hasil enkripsi dengan sha1 : %x\n", encrypted)
}
