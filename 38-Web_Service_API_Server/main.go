package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// membuat struct
type mahasiswa struct {
	ID    string
	Nama  string
	Nilai float64
}

// membuat slice berisi beberapa data mahasiswa bertipe struct mahasiswa
var dataMahasiswa = []mahasiswa{
	{"U001", "Ahmad Sidik Rudni", 3.99},
	{"U002", "Ahmad Alfian Purba", 3.98},
	{"U003", "Galuh Nugraha", 3.97},
}

// Struct mahasiswa di atas digunakan sebagai tipe elemen slice sample data, ditampung variabel dataMahasiswa

// membuat fungsi untuk menghandle endpoint /users
func users(w http.ResponseWriter, r *http.Request) {

	// menentukan tipe response, yaitu sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// mendeteksi jenis request
	if r.Method == "GET" {
		// apabila jenis requestnya adalah GET maka jalankan kode yg ada di dalam block ini

		// konversi object dataMahasiswa menjadi json
		jsonMahasiswa, err := json.Marshal(dataMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // response di-set sebagai error
			return
		}

		// mendaftarkan jsonMahasiswa sebagai response
		w.Write(jsonMahasiswa)
		return // return digunakan untuk menutup blok fungsi, sehingga baris kode program dibawahnya tidak akan diproses
	}

	// Jika hasil r.Method diatas bukan lah "GET", artinya statement if diatas tidak akan diproses, dan kode dibawah ini yang akan di proses
	http.Error(w, "", http.StatusBadRequest) // response di-set sebagai error
}

// membuat fungsi untuk menghandle endpoint /user dengan parameter id
func user(w http.ResponseWriter, r *http.Request) {
	// menentukan tipe response, yaitu sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// mengambil data form yang dikirim dari client, pada konteks ini data yang dimaksud adalah ID
		// id := r.FormValue("id")
		id := r.FormValue("id")

		// melakukan perulangan pada slice dataMahasiswa
		for _, mahasiswa := range dataMahasiswa {

			// mencari mahasiswa yang memiliki id yang sama seperti id yang diminta pada request
			if mahasiswa.ID == id {
				// parsing data mahasiswa yang ditemukan menjadi JSON
				jsonMahasiswa, err := json.Marshal(mahasiswa)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError) // jika error saat mengambil/memparsing mahasiswa, tampilkan pesan internal server error
					return
				}

				w.Write(jsonMahasiswa)
				return
			}
		}

		http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound) // jika error saat mencari mahasiswa dengan id yang sesuai, tampilkan error "Mahasiswa tidak ditemukan"
		return
	}

	http.Error(w, "", http.StatusBadRequest) // jika method yang digunakan pada request bukanlah "GET", tampilkan error StatusBadRequest
}

func main() {

	// Routing
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Web Service API berjalan di http://localhost:4135")

	// menjalankan web server
	http.ListenAndServe(":4135", nil)
}
