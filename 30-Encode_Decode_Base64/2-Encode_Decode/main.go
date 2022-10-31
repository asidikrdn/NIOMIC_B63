package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Ahmad Sidik Rudini"

	// === ENCODE ===
	// menyiapkan variabel penampung hasil encoding
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	/*
		Fungsi base64.StdEncoding.EncodedLen(len(data)) menghasilkan informasi lebar variable data ketika sudah di-encode.
		Nilai tersebut kemudian ditentukan sebagai lebar alokasi tipe []byte pada variabel encoded yang nantinya digunakan untuk menampung hasil encoding.
	*/
	// proses encoding
	// base64.StdEncoding.Encode(varPenampungHasilEncode, []byte(stringYangInginDiEncode))
	base64.StdEncoding.Encode(encoded, []byte(data))
	// meng-casting hasil encoded dari byte menjadi string
	encodedStr := string(encoded)
	// menampilkan hasil encode versi byte dan versi string
	fmt.Println("Hasil encoded asli : ", encoded)
	fmt.Println("Hasil encoded setelah di-casting ke string : ", encodedStr)

	fmt.Println()

	// === DECODE ===
	// menyiapkan variabel penampung hasil decoding
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	// pada proses decoding
	// varPanjangSliceHasilDecode, varError := base64.StdEncoding.Decode(varPenampungHasilDecode, stringEncodedYangInginDiDecode)
	a, err := base64.StdEncoding.Decode(decoded, encoded)
	/*
		fungsi Decode akan mengembalikan 2 nilai, yaitu jumlah element slice byte hasil decoding, dan error.
		sedangkan hasil decodingnya sendiri, akan disimpan pada argument pertama
	*/
	if err != nil {
		fmt.Println(err.Error())
	}
	// meng-casting hasil decoded dari byte menjadi string
	decodedStr := string(decoded)
	// menampilkan hasil decode versi byte dan versi string
	fmt.Println(a) // 18
	fmt.Println("Hasil decoded asli : ", decoded)
	fmt.Println("Hasil decoded setelah di-casting ke string : ", decodedStr)
}
