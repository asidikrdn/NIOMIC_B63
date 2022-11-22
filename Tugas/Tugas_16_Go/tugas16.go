package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mahasiswa struct {
	nim          int
	nama         string
	umur         int
	asal_kota    string
	asal_sekolah string
	nilai_ujian  float64
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_tugas16")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ambilDataMahasiswa() {
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
		err = rows.Scan(&mhs.nim, &mhs.nama, &mhs.umur, &mhs.asal_kota, &mhs.asal_sekolah, &mhs.nilai_ujian)
		if err != nil {
			panic(err.Error())
		}

		dataMahasiswa = append(dataMahasiswa, mhs)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("===================================")
	for i, mhs := range dataMahasiswa {
		fmt.Printf("Data Mahasiswa ke-%d\n", i+1)
		fmt.Println("----------------")
		fmt.Println("NIM :", mhs.nim)
		fmt.Println("Nama :", mhs.nama)
		fmt.Println("Umur :", mhs.umur)
		fmt.Println("Asal Kota :", mhs.asal_kota)
		fmt.Println("Asal Sekolah :", mhs.asal_sekolah)
		fmt.Println("===================================")
	}
}

func main() {
	ambilDataMahasiswa()
}
