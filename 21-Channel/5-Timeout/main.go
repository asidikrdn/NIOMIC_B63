package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	pesan := make(chan int)

	go kirimData(pesan)

	tampilkanPesan((pesan))
}

// Membuat fungsi mengirim data ke channel pada tiap durasi tertentu dengan durasi interval acak
func kirimData(cH chan<- int) {
	i := 1
	// membuat perulangan tak hingga dengan format do-while tanpa 'while'nya
	for {
		cH <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
		i++
	}
}

// Membuat fungsi menampilkan data dari channel menggunakan 'select'
func tampilkanPesan(cH <-chan int) {
	// fungsi tak hingga seperti berikut harus diberikan nama agar bisa dilakukan break
perulanganTerimaData:
	for {
		select {
		case data := <-cH: // akan terpenuhi ketika ada serah terima data pada channel
			fmt.Printf("Menerima pesan ke-%d \n", data)
		case <-time.After(time.Second * 5): // akan terpenuhi ketika tidak ada aktivitas penerimaan data dari channel dalam durasi 5 detik.
			fmt.Printf("Timeout, tidak ada aktivitas dalam 5 detik")
			break perulanganTerimaData
		}
	}
}
