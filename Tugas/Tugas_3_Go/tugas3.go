package main

import "fmt"

func main() {
	buah := []string{"apel", "jeruk", "anggur", "melon"}

	buah = append(buah, "pepaya")

	fmt.Println("Jumlah Element :", len(buah))

	fmt.Println("Isi Element :", buah)

	for i, namaBuah := range buah {
		fmt.Printf("Element ke - : %d %s\n", i, namaBuah)
	}
}
