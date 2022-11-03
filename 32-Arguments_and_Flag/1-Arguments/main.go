package main

import (
	"fmt"
	"os"
)

func main() {
	// mengambil argument dan menyimpannya ke dalam variabel
	argsRaw := os.Args
	fmt.Println(argsRaw) // []string{".../1-Arguments", "arg1", "arg2", "arg3"}

	// mengambil nilai argumentnya saja, tanpa melibatkan pathnya yang berada pada element ke-0
	args := argsRaw[1:]
	fmt.Println(args) // []string{"arg1", "arg2", "arg3"}

	// Jika ingin melihat atau menggunakan argument, saat eksekusi kode tuliskan argument setelah perintah eksekusi
	// Format : go run pathFile (tulis argument disini)
	// Contoh : go run 32-Argument_and_Flag/1-Arguments/main.go banana potato "ice cream"
}
