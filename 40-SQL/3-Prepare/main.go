package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mahasiswa struct {
	nim   string
	nama  string
	umur  int
	nilai float64
}

// membuat fungsi untuk koneksi ke database
func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ambilDataMahasiswa() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	querySeluruhBarisData, err := db.Prepare("select * from tb_mahasiswa")
	if err != nil {
		panic(err.Error())
	}

	rows, err := querySeluruhBarisData.Query()
	if err != nil {
		panic(err.Error())
	}

	var dataMahasiswa []mahasiswa

	for rows.Next() {
		var mhs mahasiswa
		err = rows.Scan(&mhs.nim, &mhs.nama, &mhs.umur, &mhs.nilai)
		if err != nil {
			panic(err.Error())
		}
		dataMahasiswa = append(dataMahasiswa, mhs)
	}

	fmt.Println(dataMahasiswa)
}

func ambilDataSatuPerSatuMahasiswa() {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	querySatuBarisData, err := db.Prepare("select * from tb_mahasiswa where id = ?")
	if err != nil {
		panic(err.Error())
	}

	var mhs mahasiswa

	querySatuBarisData.QueryRow("U001").Scan(&mhs.nim, &mhs.nama, &mhs.umur, &mhs.nilai)
	fmt.Println(mhs)

	querySatuBarisData.QueryRow("U002").Scan(&mhs.nim, &mhs.nama, &mhs.umur, &mhs.nilai)
	fmt.Println(mhs)

	querySatuBarisData.QueryRow("U003").Scan(&mhs.nim, &mhs.nama, &mhs.umur, &mhs.nilai)
	fmt.Println(mhs)
}

func main() {
	ambilDataMahasiswa()

	fmt.Println()

	ambilDataSatuPerSatuMahasiswa()
}
