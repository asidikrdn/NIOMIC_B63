package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// Membuat sebuah slice yang beriskan beberapa object struct
	arrMahasiswa := []User{{"Ahmad Sidik Rudini", 23}, {"M Rizki Syabani", 15}}

	// proses encoding ke JSON
	jsonData, err := json.Marshal(arrMahasiswa)
	if err != nil {
		panic(err.Error())
	}

	// casting hasil encode ke bentuk string
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	fmt.Println(arrMahasiswa[0])

	fmt.Println("=======================")
	fmt.Println()

	// Membuat sebuah object map[string]interface{}
	arrMahasiswa2 := map[string]interface{}{"Name": "Ahmad Sidik Rudini", "Age": 23}

	// proses encoding ke JSON
	jsonData2, err2 := json.Marshal(arrMahasiswa2)
	if err2 != nil {
		panic(err2.Error())
	}

	// casting hasil encode ke bentuk string
	jsonString2 := string(jsonData2)
	fmt.Println(jsonString2)

	fmt.Println(arrMahasiswa2)

	fmt.Println("=======================")
	fmt.Println()

	// Membuat sebuah slice bertipe object map[string]interface{}
	arrMahasiswa3 := []map[string]interface{}{{"Name": "Ahmad Sidik Rudini", "Age": 23}, {"Name": "M Rizki Syabani", "Age": 15}}

	// proses encoding ke JSON
	jsonData3, err3 := json.Marshal(arrMahasiswa3)
	if err3 != nil {
		panic(err2.Error())
	}

	// casting hasil encode ke bentuk string
	jsonString3 := string(jsonData3)
	fmt.Println(jsonString3)

	fmt.Println(arrMahasiswa3)
}
