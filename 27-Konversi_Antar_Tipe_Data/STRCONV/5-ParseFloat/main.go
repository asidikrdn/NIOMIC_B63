package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi string ke numerik desimal dengan lebar data tertentu.
	// varHasil, varError := strconv.ParseFloat(stringYangInginDiKonversi, lebarData)

	str := "24.12"

	// konversi string '24.12' ke bilangan desimal dengan lebar data float32
	num, err := strconv.ParseFloat(str, 32)
	if err == nil {
		fmt.Println(num) // 24.120008392334
	}
}
