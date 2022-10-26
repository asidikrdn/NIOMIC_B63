package main

import (
	"fmt"
	"time"
)

func main() {
	cH := make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		// setelah 4 detik, jalankan kode dibawah ini
		fmt.Println("expired")
		cH <- true
	})

	fmt.Println("start")
	<-cH // proses ini akan menahan kode dibawahnya sampai dengan data dari channel cH diterima
	fmt.Println("finish")
}
