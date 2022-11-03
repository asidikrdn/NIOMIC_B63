package main

import (
	"flag"
	"fmt"
)

func main() {
	// Format pengambilan flag : flag.TipeData(key, defaultValue, hint)

	// mengambil flag dengan key "name"
	name := flag.String("name", "anonymous", "tuliskan nama anda")
	// mengambil flag dengan key "age"
	age := flag.Int("age", 23, "tuliskan umur anda")

	// Mengambil nilai flag dari arg
	flag.Parse()

	// Menampilkan flag yang sudah disimpan ke dalam variabel
	fmt.Printf("Nama\t : %s\n", *name)
	fmt.Printf("Usia\t : %d\n", *age)

	// Jika ingin melihat atau menggunakan argument, saat eksekusi kode tuliskan argument setelah perintah eksekusi
	// Format : go run pathFile -(key)=(value)
	// Contoh : go run 32-Argument_and_Flag/2-Flag/main.go -name="Ahmad Sidik Rudini" -age=22
}
