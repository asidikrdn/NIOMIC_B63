package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8) // digunakan untuk menentukan jumlah prosesor	yang aktif

	/**
	Berikut merupakan contoh implementasi sederhana tentang goroutine. Program di bawah ini
	menampilkan 10 baris teks, 5 dieksekusi dengan cara biasa, dan 5 lainnya dieksekusi
	sebagai goroutine baru.
	*/
	go tampilkanPesan(5, "saya kirim")                   // eksekusi sebagai goroutine baru
	tampilkanPesan(5, "saya keduaaaaaaaaaaaaaaaaaaaaaa") // eksekusi dengan cara biasa

	var input string
	fmt.Scanln(&input)
	/**
	Fungsi fmt.Scanln() mengakibatkan proses jalannya aplikasi berhenti di baris itu (blocking)
	hingga user menekan tombol enter. Hal ini perlu dilakukan karena ada kemungkinan waktu selesainya
	eksekusi goroutine tampilkanPesan() lebih lama dibanding waktu selesainya goroutine utama main(),
	mengingat bahwa keduanya sama-sama asnychronous. Jika itu terjadi, goroutine yang belum selesai
	secara paksa dihentikan	prosesnya karena goroutine utama sudah selesai dijalankan.
	*/
}

// Membuat sebuah fungsi
func tampilkanPesan(x int, pesan string) {
	for i := 0; i < x; i++ {
		fmt.Println((i + 1), pesan)
	}
}
