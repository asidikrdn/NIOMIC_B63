package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Nama : Ahmad Sidik Rudini")

	time.Sleep(5 * time.Second)
	fmt.Println("Hobi : Travelling")

	cH := make(chan bool)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Tanggal Lahir : 14 Maret")
		cH <- true
	})
	<-cH
}
