package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi data float64 ke string dengan format eksponen, lebar digit desimal, dan lebar tipe data tertentu.
	// varHasil := strconv.FormatFloat(bilYangInginDiKonversi, 'formatEksponen', jumlahDigit, lebarData)

	num := float64(24.12)

	// konversi bilangan '24.12' ke string berbentuk desimal dengan 5 digit di belakang koma, tanpa eksponen, dan lebar data float64
	str1 := strconv.FormatFloat(num, 'f', 5, 64)
	fmt.Println(str1) // 24.12000

	// konversi bilangan '24.12' ke string berbentuk desimal dengan 5 digit di belakang koma, dengan eksponen e, dan lebar data float64
	str2 := strconv.FormatFloat(num, 'e', 5, 64)
	fmt.Println(str2) // 2.41200e+01

	/*
		Beberapa format eksponen yang dapat digunakan :
		- b : -ddddp±ddd, a, eksponen biner (basis 2)
		- e : -d.dddde±dd, a, eksponen desimal (basis 10)
		- E : -d.ddddE±dd, a, eksponen desimal (basis 10)
		- f : -ddd.dddd, tanpa eksponen
		- g : Akan menggunakan format eksponen e untuk eksponen besar dan f untuk selainnya
		- G : Akan menggunakan format eksponen E untuk eksponen besar dan f untuk selainnya
	*/
}
