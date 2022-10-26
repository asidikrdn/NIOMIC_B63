package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// Membuat slice yang berisi sekumpulann bilangan
	bilangan := []int{3, 5, 4, 2, 6, 7, 8, 10, 9, 1}
	fmt.Println("Kumpulan bilangan :", bilangan)

	// Membuat channel1 untuk menampung rata-rata
	chAverage := make(chan float64)

	// Membuat channel2 untuk menampung nilai tertinggi
	chMaxValue := make(chan int)

	go getMaxValue(bilangan, chMaxValue)
	go getRataRata(bilangan, chAverage)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-chAverage:
			fmt.Printf("Nilai rata-ratanya adalah \t : %.2f \n", avg)
		case max := <-chMaxValue:
			fmt.Printf("Nilai tertingginya adalah \t : %d \n", max)
		}
	}
}

// Membuat fungsi mencari rata-rata dengan input 2 argument,
// arg1: sekumpulan angka, arg2: channel untuk menyimpan hasil
func getRataRata(arrBil []int, cH chan float64) {
	total := 0

	// menjumlahkan seluruh bilangan anggota arrBil dengan perulangan for range
	for _, bil := range arrBil {
		// menyimpan nilainya ke variabel total
		total += bil
	}

	//  menghitung rata-rata (totalBilangan / banyakBilangan) dan dikirim ke channel yg digunakan pada arg2
	cH <- float64(total) / float64(len(arrBil))
}

// Membuat fungsi mencari rata-rata dengan input 2 argument,
// arg1: sekumpulan angka, arg2: channel untuk menyimpan hasil
func getMaxValue(arrBil []int, cH chan int) {
	max := arrBil[0]

	// mencari nilai tertinggi pada arrBil dan menyimpannya ke variabel max
	for _, bil := range arrBil {
		if max < bil {
			max = bil
		}
	}

	cH <- max
}
