package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// Membuat channel
	chPesan := make(chan string)

	// menjalankan fungsi kirimPesan dalam goroutine baru secara async
	go kirimPesan(chPesan)

	// menjalankan fungsi tampilkanPesan dalam goroutine main secara synchronous
	tampilkanPesan(chPesan)
}

// Membuat fungsi untuk mengirim data ke channel sebanyak 20x
func kirimPesan(cH chan<- string) { // Parameter `cH` hanya bisa digunakan untuk mengirim data
	for i := 0; i < 20; i++ {
		cH <- fmt.Sprintf("Data %d", i+1)
	}
	fmt.Println("Channel akan ditutup")
	close(cH) // menutup channel setelah mengirim data sebanyak 20x, agar proses penerimaan data dari channel juga turut terhenti

}

// Membuat fungsi untuk mencetak data dari channel dengan perulangan for - range
func tampilkanPesan(cH <-chan string) { // Parameter `cH` hanya bisa digunakan untuk menerima data
	// untuk perulangan for-range pada channel, tidak perlu menggunakan loadash( _ ) untuk pengganti variabel iterasi (index)
	for pesan := range cH {
		fmt.Println(pesan)
	}
}
