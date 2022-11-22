package main

import (
	"errors"
	"fmt"
)

func main() {
	var angka int

	fmt.Println("PROGRAM MEMERIKSA BILANGAN GENAP")

	defer fmt.Println("PROGRAM BERAKHIR")

	fmt.Println()

	fmt.Print("Masukkan bilangan genap : ")
	fmt.Scanln(&angka)

	if valid, err := validasi(angka); valid {
		fmt.Println(angka, "merupakan bilangan genap")
	} else {
		panic(err.Error())
	}

	fmt.Println()

}

func validasi(bil int) (bool, error) {
	if bil%2 == 0 {
		return true, nil
	} else {
		return false, errors.New("Yang anda input bukan bilangan genap")
	}
}
