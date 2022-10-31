package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi bool ke string
	// varHasil := strconv.FormatBool(nilaiBoolYangInginDiKonversi)

	bul := true

	str := strconv.FormatBool(bul)
	fmt.Println(str) // true
}
