package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Konversi string berbentuk number dengan basis tertentu ke dalam bilangan int lebar data yang dapat ditentukan
	// varPenampungHasil, varPenampungError := strconv.ParseInt(stringYangInginDiKonversi, basisBilangan, lebarData)

	bilString1 := "123"
	fmt.Println(bilString1, reflect.ValueOf(bilString1).Type()) // 123
	// konversi string number basis 10 (bilangan bulat desimal) ke bilangan int dengan lebar data int64
	bilInt1, err := strconv.ParseInt(bilString1, 10, 64)
	if err == nil {
		fmt.Println(bilInt1, reflect.ValueOf(bilInt1).Type()) // 123
	}

	bilString2 := "1011"
	fmt.Println(bilString2, reflect.ValueOf(bilString2).Type()) // 1011
	// konversi string number basis 2 (bilangan biner) ke bilangan int dengan lebar data int8
	bilInt2, err := strconv.ParseInt(bilString2, 2, 8)
	if err == nil {
		fmt.Println(bilInt2, reflect.ValueOf(bilInt2).Type()) // 11
	}
}
