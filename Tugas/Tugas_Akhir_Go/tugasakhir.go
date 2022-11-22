package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

var baseURL = "http://localhost:4135"

func fetchMahasiswa() []mahasiswa {
	var err error
	var client = &http.Client{}
	var dataMahasiswa []mahasiswa

	req, err := http.NewRequest("GET", baseURL+"/mahasiswa", nil)
	if err != nil {
		panic(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&dataMahasiswa)
	if err != nil {
		panic(err.Error())
	}

	return dataMahasiswa
}

func fetchSatuMahasiswa(nim string) mahasiswa {
	var err error
	var client = &http.Client{}
	var dataMahasiswa mahasiswa

	param := url.Values{}
	param.Set("nim", nim)
	payload := bytes.NewBufferString(param.Encode())

	req, err := http.NewRequest("GET", baseURL+"/mhs?"+payload.String(), nil)
	if err != nil {
		panic(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&dataMahasiswa)
	if err != nil {
		panic(err.Error())
	}

	return dataMahasiswa
}

func ubahMahasiswa(nim string, namaBaru string) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("update tbl_mahasiswa set nama=? where nim=?", namaBaru, nim)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Nama mahasiswa dengan nim:%s berhasil diedit\n", nim)
}

func hapusMahasiswa(nim string) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM tbl_mahasiswa WHERE nim=?", nim)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Data mahasiswa dengan nim:%s berhasil dihapus\n", nim)
}

func main() {
	// Ambil data mahasiswa
	dataMahasiswa := fetchMahasiswa()
	fmt.Println("----------")
	for i, mhs := range dataMahasiswa {
		fmt.Printf("Mahasiswa ke-%d\n", i+1)
		fmt.Println("NIM :", mhs.Nim, "\nNama :", mhs.Nama, "\nUmur :", mhs.Umur, "\nAsal Kota :", mhs.Asal_kota, "\nAsal Sekolah :", mhs.Asal_sekolah, "\nNilai Ujian :", mhs.Nilai_ujian)
		fmt.Println("----------")
	}

	fmt.Println()
	fmt.Println("======================")
	fmt.Println()

	// Ambil satu data mahasiswa
	dataMhs := fetchSatuMahasiswa("201106041165")
	fmt.Println("NIM :", dataMhs.Nim, "\nNama :", dataMhs.Nama, "\nUmur :", dataMhs.Umur, "\nAsal Kota :", dataMhs.Asal_kota, "\nAsal Sekolah :", dataMhs.Asal_sekolah, "\nNilai Ujian :", dataMhs.Nilai_ujian)

	fmt.Println()
	fmt.Println("======================")
	fmt.Println()

	// Mengubah nama mahasiswa
	ubahMahasiswa("201106041165", "Ahmad Sidik Rudini")
	fmt.Println()
	// melihat hasil ubahan
	dataMhs = fetchSatuMahasiswa("201106041165")
	fmt.Println("NIM :", dataMhs.Nim, "\nNama :", dataMhs.Nama, "\nUmur :", dataMhs.Umur, "\nAsal Kota :", dataMhs.Asal_kota, "\nAsal Sekolah :", dataMhs.Asal_sekolah, "\nNilai Ujian :", dataMhs.Nilai_ujian)

	fmt.Println()
	fmt.Println("======================")
	fmt.Println()

	// Menghapus satu mahasiswa
	hapusMahasiswa("201106041166")
	fmt.Println()
	// melihat hasil ubahan
	dataMahasiswa = fetchMahasiswa()
	fmt.Println("----------")
	for i, mhs := range dataMahasiswa {
		fmt.Printf("Mahasiswa ke-%d\n", i+1)
		fmt.Println("NIM :", mhs.Nim, "\nNama :", mhs.Nama, "\nUmur :", mhs.Umur, "\nAsal Kota :", mhs.Asal_kota, "\nAsal Sekolah :", mhs.Asal_sekolah, "\nNilai Ujian :", mhs.Nilai_ujian)
		fmt.Println("----------")
	}
}
