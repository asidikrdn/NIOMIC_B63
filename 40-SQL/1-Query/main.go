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

// membuat fungsi untuk mengambil beberapa baris data dari database dengan Query()
func ambilDataMahasiswa() {
	// menghubungkan database
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close() // menutup database ketika sudah tidak digunakan

	// mengambil data dari database dengan Query()
	rows, err := db.Query("select * from tb_mahasiswa")
	if err != nil {
		panic(err)
	}
	defer rows.Close() // menutup instance koneksi

	// menyiapkan variabel penampung hasil query
	var dataMahasiswa []mahasiswa

	// menyimpan hasil query baris per baris dengan perulangan
	for rows.Next() {
		// tiap baris data akan disimpan sementara ke variabel mhs
		var mhs = mahasiswa{}
		err := rows.Scan(&mhs.id, &mhs.name, &mhs.age, &mhs.grade)
		if err != nil {
			panic(err.Error())
		}
		// satu per satu data tiap baris dikirim ke slice dataMahasiswa
		dataMahasiswa = append(dataMahasiswa, mhs)
	}

	// jika terjadi error saat mengambil data dari database
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	// menampilkan tiap baris data mahasiswa
	for i, mhs := range dataMahasiswa {
		fmt.Printf("Mahasiswa ke-%d\n", i+1)
		fmt.Println("NIM :", mhs.id, "\nNama :", mhs.name, "\nUmur :", mhs.age, "\nNilai :", mhs.grade)
		fmt.Println()
	}

}

// membuat fungsi untuk mengambil satu baris data dari database dengan Query() dan id sebagai parameternya
func ambilDataSatuMahasiswa(id string) {
	// menghubungkan database
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close() // menutup database ketika sudah tidak digunakan

	// mengambil data dari database dengan Query()
	rows, err := db.Query("select name, age, grade from tb_mahasiswa where id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close() // menutup instance koneksi

	// menyiapkan variabel penampung hasil query
	var dataMahasiswa []mahasiswa

	// menyimpan hasil query baris per baris dengan perulangan
	for rows.Next() {
		// tiap baris data akan disimpan sementara ke variabel mhs
		var mhs = mahasiswa{}
		err := rows.Scan(&mhs.name, &mhs.age, &mhs.grade)
		if err != nil {
			panic(err.Error())
		}
		// satu per satu data tiap baris dikirim ke slice dataMahasiswa
		dataMahasiswa = append(dataMahasiswa, mhs)
	}

	// jika terjadi error saat mengambil data dari database
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	// menampilkan tiap baris data mahasiswa, namun karena hanya ada 1 baris data, maka hanya akan tampil 1 baris data saja
	for _, mhs := range dataMahasiswa {
		fmt.Println("Nama :", mhs.name, "\nUmur :", mhs.age, "\nNilai :", mhs.grade)
	}

}

func main() {
	fmt.Println("--------------------------------")
	ambilDataMahasiswa()
	fmt.Println("--------------------------------")

	fmt.Println()

	fmt.Println("Ambil satu mahasiswa yang memiliki ID U001")
	ambilDataSatuMahasiswa("U001")
}
