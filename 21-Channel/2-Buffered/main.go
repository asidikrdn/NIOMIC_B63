package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// Membuat channel pesan dengan ukuran buffernya adalah 2 (0,1,2)
	pesan := make(chan int, 2)

	// Membuat goroutine dengan sebuah anonymous func yang langsung dieksekusi
	go func() {
		for {
			i := <-pesan // mengambil nilai dalam channel pesan dan disimpan pada variabel i
			fmt.Println("Menerima pesan", i)
		}
	}()

	for i := 0; i < 6; i++ {
		fmt.Println("Mengirim pesan", i)
		pesan <- i // mengirim nilai variabel i ke channel pesan
	}

	/**
	Bisa dilihat pada outputnya di terminal, proses pengiriman data akan berhenti di pesan ke 3, karena ukuran buffer penuh.
	Setelah isi buffer dikosongkan, barulah pesan 4 dan 5 bisa dikirim.
	*/

	// Untuk menahan agar terminal program tidak terhenti
	var input string
	fmt.Scanln(&input)
}
