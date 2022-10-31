package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	bilInt := 1234

	fmt.Println(bilInt, reflect.ValueOf(bilInt).Type())

	// Konversi Int ke String
	// varPenampungHasil := strconv.Itoa(bilanganYangInginDiKonversi)
	bilString := strconv.Itoa(bilInt)

	fmt.Println(bilString, reflect.ValueOf(bilString).Type())
}
