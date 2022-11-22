package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type mahasiswa struct {
	Nim          string  `json:"nim"`
	Nama         string  `json:"nama"`
	Umur         int     `json:"umur"`
	Asal_kota    string  `json:"asal_kota"`
	Asal_sekolah string  `json:"asal_sekolah"`
	Nilai_ujian  float64 `json:"nilai_ujian"`
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_tugas16")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ambilDataMahasiswa() []mahasiswa {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_mahasiswa")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var dataMahasiswa []mahasiswa
	for rows.Next() {
		var mhs mahasiswa
		err = rows.Scan(&mhs.Nim, &mhs.Nama, &mhs.Umur, &mhs.Asal_kota, &mhs.Asal_sekolah, &mhs.Nilai_ujian)
		if err != nil {
			panic(err.Error())
		}
		dataMahasiswa = append(dataMahasiswa, mhs)
	}

	return dataMahasiswa
}

func ambilDataSatuMahasiswa(nim string) mahasiswa {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var mhs mahasiswa

	err = db.QueryRow("SELECT * FROM tbl_mahasiswa WHERE nim = ?", nim).Scan(&mhs.Nim, &mhs.Nama, &mhs.Umur, &mhs.Asal_kota, &mhs.Asal_sekolah, &mhs.Nilai_ujian)
	if err != nil {
		panic(err.Error())
	}

	return mhs
}

func main() {
	http.HandleFunc("/mahasiswa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dataMahasiswa []mahasiswa = ambilDataMahasiswa()

		if r.Method == "GET" {
			jsonMahasiswa, err := json.Marshal(dataMahasiswa)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(jsonMahasiswa)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	})

	http.HandleFunc("/mhs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		nim := r.FormValue("nim")

		var dataMahasiswa mahasiswa = ambilDataSatuMahasiswa(nim)

		if r.Method == "GET" {
			jsonMahasiswa, err := json.Marshal(dataMahasiswa)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(jsonMahasiswa)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	})

	fmt.Println("Web Server Running on http://localhost:4135")
	http.ListenAndServe(":4135", nil)
}
