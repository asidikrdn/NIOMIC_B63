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
	} else if r.Method == "POST" {
		// jika methodnya adalah POST, maka tambahkan data berikut

		// membuat variabell penampung request body
		var dataMahasiswaBaru mahasiswa

		// decoding request body ke variabel dataMahasiswaBaru yang bertipe struct mahasiswa
		err := json.NewDecoder(r.Body).Decode(&dataMahasiswaBaru)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(dataMahasiswaBaru)

		// menambahkan object dataMahasiswaBaru ke dalam slice dataMahasiswa
		dataMahasiswa = append(dataMahasiswa, dataMahasiswaBaru)

		x, _ := json.Marshal("Data Berhasil Ditambahkan !")
		w.Write(x)
		return
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
	} else if r.Method == "PATCH" {
		// jika methodnya adalah PATCH, maka edit data mahasiswa yang sesuai dengan id yang di request

		// mengambil data form yang dikirim dari client, pada konteks ini data yang dimaksud adalah ID
		id := r.FormValue("id")

		// membuat variabel untuk menampung request body
		var mahasiswaEdited mahasiswa

		// membuat variabell penampung data mahasiswa sementara
		var dataMahasiswaBaru []mahasiswa

		// decoding request body ke variabel mahasiswaEdited yang bertipe struct mahasiswa
		err := json.NewDecoder(r.Body).Decode(&mahasiswaEdited)
		if err != nil {
			panic(err.Error())
		}

		// melakukan perulangan untuk memfilter slice dataMahasiswa
		for _, mhs := range dataMahasiswa {

			if mhs.ID != id {
				// jika mahasiswa memiliki id selain yang diminta pada request, tambahkan data tsb ke dataMahasiswaBaru
				dataMahasiswaBaru = append(dataMahasiswaBaru, mhs)
			} else if mhs.ID == id {
				// jika mahasiswa memiliki id seperti yang diminta pada request, edit datanya dengan data request Body, lalu tambahkan data tsb ke dataMahasiswaBaru
				mhs.ID = mahasiswaEdited.ID
				mhs.Nama = mahasiswaEdited.Nama
				mhs.Nilai = mahasiswaEdited.Nilai

				dataMahasiswaBaru = append(dataMahasiswaBaru, mhs)
			}
		}

		// mengganti isi dataMahasiswa dengan data mahasiswa yang sudah di edit
		dataMahasiswa = dataMahasiswaBaru

		x, _ := json.Marshal("Data Berhasil Diedit!")
		w.Write(x)

		// http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound) // jika error saat mencari mahasiswa dengan id yang sesuai, tampilkan error "Mahasiswa tidak ditemukan"
		return
	} else if r.Method == "DELETE" {
		// jika methodnya adalah DELETE, maka hapus data berikut

		// mengambil data form yang dikirim dari client, pada konteks ini data yang dimaksud adalah ID
		id := r.FormValue("id")

		// membuat variabel untuk menampung data mahasiswa yang telah difilter
		var mahasiswaFiltered []mahasiswa

		// melakukan perulangan untuk memfilter slice dataMahasiswa
		for _, mhs := range dataMahasiswa {

			// mencari mahasiswa yang memiliki id selain yang diminta pada request
			if mhs.ID != id {
				mahasiswaFiltered = append(mahasiswaFiltered, mhs)
			}
		}

		// mengganti isi dataMahasiswa dengan data mahasiswa yang sudah difilter
		dataMahasiswa = mahasiswaFiltered

		x, _ := json.Marshal("Data Berhasil Dihapus!")
		w.Write(x)

		// http.Error(w, "Mahasiswa tidak ditemukan", http.StatusNotFound) // jika error saat mencari mahasiswa dengan id yang sesuai, tampilkan error "Mahasiswa tidak ditemukan"
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
