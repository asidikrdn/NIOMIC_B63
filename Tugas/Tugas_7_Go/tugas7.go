package main

import (
	"fmt"
	"reflect"
	"runtime"
)

var bilangan int = 5
var text string = "Hello"

func bacaTipeData1() {
	fmt.Println("variabel bilangan dengan nilai", bilangan, "tipe datanya :", reflect.ValueOf(bilangan).Type())
	// fmt.Println("variabel text berisi", text, "tipe datanya :", reflect.ValueOf(text).Type())
}

func bacaTipeData2() {
	// fmt.Println("variabel bilangan dengan nilai", bilangan, "tipe datanya :", reflect.ValueOf(bilangan).Type())
	fmt.Println("variabel text berisi", text, "tipe datanya :", reflect.ValueOf(text).Type())
}

func main() {
	runtime.GOMAXPROCS(2)

	go bacaTipeData2()
	bacaTipeData1()

	fmt.Scanln()
}
