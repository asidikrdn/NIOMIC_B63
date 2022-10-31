package main

import "fmt"

func main() {
	/*
		String sebenarnya adalah slice/array byte. Di Go sebuah karakter biasa (bukan unicode) direpresentasikan oleh sebuah elemen slice byte.
		Tiap elemen slice berisi data int dengan basis desimal, yang merupakan kode ASCII dari karakter dalam string.
	*/

	// Contoh 1, konversi string menjadi kumpulan byte
	// membuat string 'halo'
	text1 := "halo"
	// konversi string menjadi slice byte
	b := []byte(text1)
	// menampilkan isi variabel b yang berupa slice byte
	fmt.Printf("%d \n", b) // [10 97 108 111]

	fmt.Println()

	// Contoh 2, konversi byte menjadi string
	// membuat slice yang berisi kode ASCII beberapa huruf
	byte1 := []byte{104, 97, 108, 111}
	// konversi slice byte menjadi sebuah string
	s := string(byte1)
	// menampilkan isi variabel s yang beruoa string
	fmt.Printf("%s \n", s) // halo

	fmt.Println()

	// Contoh 3, konversi char menjadi byte, dan sebaliknya
	// konversi string 'h' menjadi kode ASCII
	c := int64('h')
	fmt.Printf("%d \n", c) // 104
	// konversi kode ASCII 104 menjadi string
	d := string(104)
	fmt.Printf("%s \n", d) // h

}
