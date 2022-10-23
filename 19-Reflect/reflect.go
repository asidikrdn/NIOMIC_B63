package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 10
	// var number = "a"
	var reflectnumber = reflect.ValueOf(number)
	fmt.Println("Tipe Variabel :", reflectnumber.Type())

	if reflectnumber.Kind() == reflect.Int {
		fmt.Println("Nilai Variabel :", reflectnumber.Int())
	} else {
		fmt.Println("Nilai Variabel bukan merupakan integer")
	}

	fmt.Println()

	nama := "Sidik"
	fmt.Println("Tipe Variabel 'nama' adalah :", reflect.TypeOf(nama))
	fmt.Println("Nilai Variabel 'nama' adalah :", reflect.ValueOf(nama))
	fmt.Println()
}
