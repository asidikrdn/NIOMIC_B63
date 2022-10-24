package main

import (
	"fmt"
	"runtime"
)

// Membuat channel
var pesan = make(chan string)

func main() {
	runtime.GOMAXPROCS(2)

	// Menjalankan beberapa goroutine. Goroutine akan berjalan secara async
	go hello("a")
	go hello("b")
	go hello("c")

	/**
	Penerimaan channel bersifat blocking. Artinya statement var message1 = <-pesan hingga
	setelahnya tidak akan dieksekusi sebelum ada data yang dikirim lewat channel
	Jadi, Ketiga goroutine tersebut datanya akan diterima secara berurutan oleh message1 ,
	message2 , message3 ; untuk kemudian ditampilkan.
	*/
	var message1 = <-pesan
	fmt.Println(message1)
	var message2 = <-pesan
	fmt.Println(message2)
	var message3 = <-pesan
	fmt.Println(message3)
	/**
	Akan tetapi, data yang ditampilkan bisa saja acak meskipun proses penerimaan data berjalan secara berurutan.
	Hal ini karena kita tidak tau goroutine yang mana yang diproses lebih dulu. Goroutine yang dieksekusi lebih awal, datanya akan diterima
	lebih awal.
	*/
}

// Mengirim data dari sebuah goroutine ke channel
func hello(nama string) {
	var data = fmt.Sprintf("halo %s", nama)
	pesan <- data
}
