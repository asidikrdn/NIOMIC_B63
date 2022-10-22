package main

import "fmt"

func main() {
	// Bilangan positif
	var number_positif uint8 = 200 // 0-255
	fmt.Println(number_positif)    // 200
	// number_positif = 400           // Error, cannot use 400 (untyped int constant) as uint8 value in assignment (overflows)
	// fmt.Println(number_positif)    // Error, cannot use 400 (untyped int constant) as uint8 value in assignment (overflows)

	// Bilangan negatif
	var number_negatif int64 = -99
	fmt.Println(number_negatif)

	// Bilangan desimal
	var desimal float64 = 2.55
	fmt.Println(desimal)

	// Boolean
	var apakah_ada bool = true
	fmt.Println(apakah_ada)
	apakah_ada = false
	fmt.Println(apakah_ada)

	// String
	var pesan string = "Halo Teman" // bisa menggunakan kutip dua
	fmt.Println(pesan)
	pesan = `Hai Teman` // bisa menggunakan backtick
	fmt.Println(pesan)

}
