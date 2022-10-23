# Function pada GoLang

Fungsi : Sebuah blok kode yang dibungkus dengan nama tertentu

Dari awal pembelajaran ini, kita sudah menggunakan sebuah fungsi, yaitu main.

Function main adalah fungsi utama pada GoLang yang akan dijalankan pertama kali.

Kita juga bisa membuat fungsi sendiri, pembuatan fungsi tersebut harus dilakukan diluar function main.

Ada 3 cara pembuatan fungsi pada GoLang

- Membuat fungsi tanpa return value
- Membuat fungsi dengan return value
- Membuat fungsi dengan return value dan parameter

Untuk membuat fungsi tanpa return value caranya dengan menuliskan format berikut :
`func nama_fungsi(){ ...isi dari fungsi}`

Untuk membuat fungsi tanpa return value caranya dengan menuliskan format berikut :

```Go
func nama_fungsi(){ ..isi dari fungsi}
```

Untuk membuat fungsi dengan return value caranya dengan menuliskan format berikut :

```Go
func nama_fungsi()tipe_data_dari_return_value{ 
  ..isi dari fungsi
  return return_value
}
```

Untuk membuat fungsi dengan return value dan parameter caranya dengan menuliskan format berikut :

```Go
func nama_fungsi(variabel_parameter tipe_data_parameter)tipe_data_dari_return_value{ 
  ..isi dari fungsi
  return return_value
}
```

## Fungsi dengan Multiple Return Value

Berbeda dengan bahasa pemrograman lainnya yang biasanya hanya bisa mengembalikan sebuah nilai dalam satu fungsi, pada GoLang sebuah fungsi dapat mengembalikan beberapa nilai sekaligus.

Format untuk membuat fungsi dengan beberapa nilai kembalian adalah sebagai berikut (misalnya sebuah fungsi dengan 2 nilai kembalian) :

```Go
func nama_fungsi(variabel_parameter1 tipe_data_parameter1,variabel_parameter2 tipe_data_parameter2)(tipe_data_dari_return_value1, tipe_data_dari_return_value2){ 
  ..isi dari fungsi
  return return_value1, return_value2
}
```

Dan cara menggunakan/memanggilnya adalah sebagai berikut :

```Go
var variabel_penampung_nilai_kembalian1, variabel_penampung_nilai_kembalian2 = nama_fungsi(nilai_argument1,nilai_argument2)
```
