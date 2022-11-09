package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Menyiapkan struct untuk digunakan sebagai tipe data penampung hasil query
type mahasiswa struct {
	id    string
	name  string
	age   int
	grade float64
}

// membuat fungsi untuk koneksi ke database
func connect() (*sql.DB, error) {
	// memulai koneksi ke database
	// varDatabase, varError := sql.Open(namaDriver, stringKoneksi)
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// membuat fungsi untuk mengambil 1 baris data dari database dengan QueryRow
func ambilDataSatuMahasiswa(id string) {
	// menghubungkan database
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close() // menutup database ketika sudah tidak digunakan

	var dataMahasiswa mahasiswa

	// mengambil data dengan QueryRow() dan langsung di-chain dengan method Scan()
	err = db.QueryRow("select * from tb_mahasiswa where id = ?", id).Scan(&dataMahasiswa.id, &dataMahasiswa.name, &dataMahasiswa.age, &dataMahasiswa.grade)
	if err != nil {
		panic(err)
	}

	fmt.Println("NIM :", dataMahasiswa.id, "\nNama :", dataMahasiswa.name, "\nUmur :", dataMahasiswa.age, "\nNilai :", dataMahasiswa.grade)
}

func main() {
	fmt.Println("Ambil satu mahasiswa yang memiliki ID U001")
	ambilDataSatuMahasiswa("U001")
}
