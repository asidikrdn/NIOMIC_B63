package main

import "fmt"

func main() {
	var bil1, bil2 int

	fmt.Print("Masukkan bilangan pertama : ")
	fmt.Scanln(&bil1)
	fmt.Print("Masukkan bilangan kedua : ")
	fmt.Scanln(&bil2)

	intBil1 := int(bil1)
	intBil2 := int(bil2)

	fmt.Println()

	fmt.Printf("%d + %d = %d\n", intBil1, intBil2, intBil1+intBil2)
	fmt.Printf("%d - %d = %d\n", intBil1, intBil2, intBil1-intBil2)
	fmt.Printf("%d x %d = %d\n", intBil1, intBil2, intBil1*intBil2)
	fmt.Printf("%d / %d = %d\n", intBil1, intBil2, intBil1/intBil2)

}
