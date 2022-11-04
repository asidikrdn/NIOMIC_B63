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

/*
Salah satu property-nya yaitu FullName memiliki tag json:"Name". Tag tersebut digunakan untuk mapping informasi json ke property yang bersangkutan.
Dengan menambahkan tag json, maka property FullName struct akan secara cerdas menampung data json property Name.
Satu lagi, pada kasus decoding data json string ke variabel objek struct, semua level akses property struct penampung harus publik.
*/

func main() {
	// Membuat json string
	jsonString := `{"Name" : "Ahmad Sidik Rudini", "Age" : 23}`

	// membuat data json dengan memng-casting jsonString ke byte, karena fungsi unmarshal hanya menerima data json dalam bentuk []byte
	jsonData := []byte(jsonString)

	// membuat variabel dataUser bertipekan struct User, variabel ini akan digunakan untuk menyimpan hasil decode
	var dataUser User

	// proses decode
	// Format => varError := json.Unmarshal(dataJSONYangInginDiDecode, &varPenampungHasil)
	err := json.Unmarshal(jsonData, &dataUser)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("user :", dataUser.FullName)
	fmt.Println("age :", dataUser.Age)
}
