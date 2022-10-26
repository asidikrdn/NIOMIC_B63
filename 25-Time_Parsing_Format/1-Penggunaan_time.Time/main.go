package main

import (
	"fmt"
	"time"
)

func main() {
	// menjadikan waktu sekarang sebagai time
	time1 := time.Now()
	fmt.Printf("time1 : %v \n", time1) // time1 : 2015-09-01 17:59:31.73600891 +0700 WIB

	// Membuat time dengan data ditentukan sendiri
	// time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
	time2 := time.Date(1999, 03, 14, 05, 00, 00, 00, time.UTC)
	fmt.Printf("time2 : %v \n", time2) // time2 : 2011-12-24 10:20:00 +0000 UTC
}
