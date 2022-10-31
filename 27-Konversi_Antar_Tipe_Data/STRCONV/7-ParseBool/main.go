package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi string ke bool
	// varHasil, varError := strconv.ParseBool(stringYangInginDiKonversi)

	str := "true"

	// konversi string "true" menjadi nilai bool true
	bul, err := strconv.ParseBool(str)
	if err == nil {
		fmt.Println(bul) // true
	}

	// konversi string "1" menjadi nilai bool true
	bul, err = strconv.ParseBool("1")
	if err == nil {
		fmt.Println(bul) // true
	}

	// konversi string "false" menjadi nilai bool false
	str = "false"
	bul, err = strconv.ParseBool(str)
	if err == nil {
		fmt.Println(bul) // false
	}

	// konversi string "0" menjadi nilai bool false
	bul, err = strconv.ParseBool("0")
	if err == nil {
		fmt.Println(bul) // false
	}
}
