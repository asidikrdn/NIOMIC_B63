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

## Fungsi Variadic

Golang mengadopsi konsep variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas. Maksud tak terbatas disini adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.

Parameter variadic memiliki sifat yang mirip dengan slice. Nilai parameter-parameter yang disisipkan memiliki tipe data yang sama, dan akan ditampung oleh sebuah variabel saja. Cara pengaksesan tiap datanya juga sama, dengan menggunakan indeks.

Deklarasi parameter variadic sama dengan cara deklarasi variabel biasa, pembedanya pada parameter jenis ini ditambahkan tanda 3 titik ( ... ) setelah penulisan variabel (sebelum tipe data). Nantinya semua nilai yang disisipkan sebagai parameter akan ditampung oleh variabel tersebut.

Nilai tiap parameter bisa diakses seperti cara pengaksesan tiap elemen slice. Pada contoh kode, metode yang dipilih adalah for - range . Format penggunaan for - range adalah seperti berikut :

```Go
for variabel_penampung_index, nama_array_yang_diproses := range nama_array_yang_diproses {
  fmt.Printf("bilangan ke-%d adalah %d \n", variabel_penampung_index, nama_array_yang_diproses)
 }
```

Dalam contoh kode terdapat fungsi fmt.Sprintf(). Fungsi fmt.Sprintf() pada dasarnya sama dengan fmt.Printf() , hanya saja fungsi ini tidak menampilkan nilai, melainkan mengembalikan nilainya dalam bentuk string.
