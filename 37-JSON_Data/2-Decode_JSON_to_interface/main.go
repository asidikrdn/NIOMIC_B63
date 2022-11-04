package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Membuat json string
	jsonString := `{"Name" : "Ahmad Sidik Rudini", "Age" : 23}`

	// membuat data json dengan memng-casting jsonString ke byte, karena fungsi unmarshal hanya menerima data json dalam bentuk []byte
	jsonData := []byte(jsonString)

	// membuat variabel bertipe map[string]interface{} untuk menampung hasil decode
	var data1 map[string]interface{}

	// proses decode JSON
	err := json.Unmarshal(jsonData, &data1)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("data hasil decode :", data1)
	fmt.Println("user :", data1["Name"])
	fmt.Println("age :", data1["Age"])

	fmt.Println()

	// membuat variabel bertipe map[string]interface{} untuk menampung hasil decode
	var data2 interface{}

	// proses decode JSON
	err = json.Unmarshal(jsonData, &data2)
	if err != nil {
		panic(err.Error())
	}
	decodedData := data2.(map[string]interface{}) // casting object interface{} ke map[string]interface{}

	fmt.Println("data hasil decode :", data2)
	fmt.Println("data hasil decode + casting :", decodedData)
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age :", decodedData["Age"])

	fmt.Println()
}
