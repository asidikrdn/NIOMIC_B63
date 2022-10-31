package main

import "fmt"

func main() {
	/**
	Type assertions merupakan teknik untuk mengambil tipe data konkret dari data yang terbungkus dalam interface{}.
	Jadi bisa disimpulkan bahwa teknik type assertions hanya bisa dilakukan pada data bertipe interface{}. Lebih jelasnya silakan cek contoh berikut.
	*/

	// Membuat sebuah map dengan interface kosong
	data := map[string]interface{}{
		"nama":    "Ahmad Sidik",
		"tinggi":  164.5,
		"berat":   52,
		"isMale":  true,
		"hobbies": []string{"travelling", "swimming"},
	}

	// Menampilkan isi map sambil mengasertasi ke tipe data tertentu
	// namaMap[key].(tipeDataYangDiinginkan)
	fmt.Println(data["nama"].(string))
	fmt.Println(data["tinggi"].(float64))
	fmt.Println(data["berat"].(int))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))
	/**
	Pada kode di atas, tidak akan terjadi panic error, karena semua operasi type assertion adalah dilakukan menggunakan tipe data yang sudah sesuai
	dengan tipe data nilai aslinya. Seperti data["nama"] yang merupakan string pasti bisa di-asertasi ke tipe string.
	Coba lakukan asertasi ke tipe yang tidak sesuai dengan tipe nilai aslinya, seperti data["nama"].(int), pasti akan men-trigger panic error.
	*/

	fmt.Println()

	/** Nah, dari penjelasan di atas, terlihat bahwa kita harus tau terlebih dahulu apa tipe data asli dari data yang tersimpan dalam interface.
	Jika misal tidak tau, maka bisa gunakan teknik di bawah ini untuk pengecekan sukses tidaknya proses asertasi.

	Tipe asli data pada variabel interface{} bisa diketahui dengan cara meng-casting ke tipe type, namun casting ini hanya bisa dilakukan pada switch.
	*/

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
