package main

import (
	"encoding/json"
	"fmt"
)

// membuat struct yang nantinya digunakan untuk membuat variabel baru penampung hasil decode json string
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// Membuat json string dengan beberapa data
	jsonString := `[{"Name" : "Ahmad Sidik Rudini", "Age" : 23}, {"Name" : "M Rizki Syabani", "Age" : 15}]`

	// membuat data json dengan memng-casting jsonString ke byte, karena fungsi unmarshal hanya menerima data json dalam bentuk []byte
	jsonData := []byte(jsonString)

	// membuat variabel slice dataUser bertipekan struct User, variabel ini akan digunakan untuk menyimpan hasil decode
	var dataUser []User

	// proses decode
	// Format => varError := json.Unmarshal(dataJSONYangInginDiDecode, &varPenampungHasil)
	err := json.Unmarshal(jsonData, &dataUser)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("object hasil decode :", dataUser)
	fmt.Println()
	// menampilkan data user ke 1
	fmt.Println("user :", dataUser[0].FullName)
	fmt.Println("age :", dataUser[0].Age, "th")
	fmt.Println()
	// menampilkan data user ke 2
	fmt.Println("user :", dataUser[1].FullName)
	fmt.Println("age :", dataUser[1].Age, "th")
}
