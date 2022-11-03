package main

import (
	"flag"
	"fmt"
)

func main() {
	// Format pengambilan flag : flag.`TipeData`Var(&varPenampung, key, defaultValue, hint)

	// mengambil flag dengan key "name"
	var name string
	flag.StringVar(&name, "name", "anonymous", "tuliskan nama anda")
	// mengambil flag dengan key "age"
	var age int
	flag.IntVar(&age, "age", 23, "tuliskan umur anda")

	// Mengambil nilai flag dari arg versi kedua
	flag.Parse()

	// Menampilkan flag yang sudah disimpan ke dalam variabel, dengan cara kedua kita bisa langsung memanggil variabelnya (bukan pointernya)
	fmt.Printf("Nama\t : %s\n", name)
	fmt.Printf("Usia\t : %d\n", age)

	// Jika ingin melihat atau menggunakan argument, saat eksekusi kode tuliskan argument setelah perintah eksekusi
	// Format : go run pathFile -(key)=(value)
	// Contoh : go run 32-Argument_and_Flag/2-Flag/main.go -name="Ahmad Sidik Rudini" -age=22
}
