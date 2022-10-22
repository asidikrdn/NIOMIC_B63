package main

import "fmt"

func main() {
	const x = 30 // Membuat konstanta x
	fmt.Println(x)
	// x = 25	// Error, cannot assign to x (untyped int constant 30)
	fmt.Printf("%d", x)
}
