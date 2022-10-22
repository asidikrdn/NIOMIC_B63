package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	// Penjumlahan
	fmt.Printf("Penjumlahan %d + %d = %d \n", a, b, a+b)

	// Pengurangan
	fmt.Printf("Pengurangan %d - %d = %d \n", a, b, a-b)

	// Perkalian
	fmt.Printf("Perkalian %d * %d = %d \n", a, b, a*b)

	// Pembagian
	fmt.Printf("Pembagian %d / %d = %d \n", a, b, a/b)

	// Modulus
	fmt.Printf("Modulus %d mod %d = %d", a, b, a%b)
}
