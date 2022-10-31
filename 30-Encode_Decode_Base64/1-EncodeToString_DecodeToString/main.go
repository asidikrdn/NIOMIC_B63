package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Ahmad Sidik Rudini"

	// EncodingToString
	// varHasilEncode := base64.URLEncoding.EncodeToString([]byte(stringYangInginDiEncode))
	// Note : []byte() digunakan untuk meng-casting string yang ingin diencode
	encodedStr := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Hasil encode :", encodedStr)

	fmt.Println()

	// DecodingToString
	// varHasilDecodeAsli := base64.URLEncoding.DecodeString(StringEncodedYangInginDiDecode)
	// varHasilCastingByteToString := string(varHasilDecodeAsli)
	decodedByte, _ := base64.StdEncoding.DecodeString(encodedStr)
	decodedStr := string(decodedByte)
	fmt.Println("Hasil decoded asli :", decodedByte)
	fmt.Println("Hasil decoded setelah dikonversi ke string :", decodedStr)
}
