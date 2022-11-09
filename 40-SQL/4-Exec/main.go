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

// membuat fungsi untuk koneksi database
func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// membuat fungsi untuk menambahkan data ke database
func tambahData(mhsBaru mahasiswa) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_mahasiswa values (?,?,?,?)", mhsBaru.nim, mhsBaru.nama, mhsBaru.umur, mhsBaru.nilai)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil ditambahkan")
}

// membuat fungsi untuk menghapus data dari database
func hapusData(id string) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("delete from tb_mahasiswa where id = ?", id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil dihapus")
}

// membuat fungsi untuk mengedit nama dengan id tertentu pada database
func editNama(nama string, id string) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("update tb_mahasiswa set name=? where id=?", nama, id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Data mahasiswa dengan id:%s berhasil diedit\n", id)
}

// membuat fungsi untuk mengedit umur dengan id tertentu pada database
func editUmur(umur int, id string) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("update tb_mahasiswa set age=? where id=?", umur, id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Data mahasiswa dengan id:%s berhasil diedit\n", id)

}

func main() {
	// menambahkan 3 mahasiswa baru
	tambahData(mahasiswa{"U004", "Ahmad", 23, 3.99})
	tambahData(mahasiswa{"U005", "Sidik", 23, 3.99})
	tambahData(mahasiswa{"U006", "Rudini", 23, 3.99})

	fmt.Println()

	// menghapus mahasiswa yang memiliki id U002
	hapusData("U002")

	fmt.Println()

	// mengedit nama mahasiswa U001
	editNama("A Sidik Rudini", "U001")

	fmt.Println()

	// mengedit umur mahasiswa U003
	editUmur(20, "U003")
}
