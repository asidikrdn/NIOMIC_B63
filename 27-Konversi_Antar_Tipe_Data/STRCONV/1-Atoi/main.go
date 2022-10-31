package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	bilString := "1234"

	fmt.Println(bilString, reflect.ValueOf(bilString).Type())

	// Konversi String ke Int
	// varPenampungHasil, varPenampungError := strconv.Atoi(stringYangInginDiKonversi)
	bilInt, err := strconv.Atoi(bilString)

	if err == nil {
		fmt.Println(bilInt, reflect.ValueOf(bilInt).Type())
	}
}
