package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(4 * time.Second) // setelah 4 detik program berjalan, variabel timer akan berisi sebuah object yang memiliki method C
	fmt.Println("start")
	<-timer.C // jika ada proses penerimaan data dari channel, maka kode dibawahnya tidak akan diproses sebelum data channelnya diterima
	fmt.Println("finish")
}
