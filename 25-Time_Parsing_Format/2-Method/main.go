package main

import (
	"fmt"
	"time"
)

func main() {
	hariIni := time.Now()

	fmt.Println("Sekarang tahun : ", hariIni.Year())            // 2022
	fmt.Println("Sekarang bulan : ", hariIni.Month())           // October
	fmt.Println("Sekarang hari : ", hariIni.Weekday())          // Wednesday
	fmt.Println("Sekarang hari : ", hariIni.Weekday().String()) // Wednesday
}
