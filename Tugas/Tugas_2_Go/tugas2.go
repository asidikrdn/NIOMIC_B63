package main

import "fmt"

func main() {
	var angka int

	fmt.Println("Silahkan inputkan bilangan :")
	fmt.Scanf("%d", &angka)
	fmt.Println()

	if angka%2 == 0 {
		fmt.Printf("%d merupakan bilangan genap", angka)
	} else {
		fmt.Printf("%d bukan bilangan genap", angka)
	}
}
