# Konversi Antar Tipe Data

Pada chapter sebelum-sebelumnyanya kita sudah mengaplikasikan beberapa cara konversi data, contohnya seperti konversi `string ↔ int` menggunakan `strconv`, dan `time.Time ↔ string`. Pada chapter ini kita akan belajar lebih banyak.

## Konversi Menggunakan `strconv`

Package `strconv` berisi banyak fungsi yang sangat membantu kita untuk melakukan konversi. Berikut merupakan beberapa fungsi yang dalam package tersebut.

- `strconv.Avoi()` : Fungsi ini digunakan untuk konversi data dari tipe string ke int. strconv.Atoi() menghasilkan 2 buah nilai kembalian, yaitu hasil konversi dan error (jika konversi sukses, maka error berisi nil).
- `strconv.Itoa()` : Merupakan kebalikan dari strconv.Atoi, berguna untuk konversi int ke string.
- `strconv.ParseInt()` : Digunakan untuk konversi string berbentuk numerik dengan basis tertentu ke tipe numerik non-desimal dengan lebar data bisa ditentukan.
- `strconv.FormatInt()` : Berguna untuk konversi data numerik int64 ke string dengan basis numerik bisa ditentukan sendiri.
- `strconv.ParseFloat()` : Digunakan untuk konversi string ke numerik desimal dengan lebar data bisa ditentukan.
- `strconv.FormatFloat()` : Berguna untuk konversi data bertipe float64 ke string dengan format eksponen, lebar digit desimal, dan lebar tipe data bisa ditentukan.
- `strconv.ParseBool()` : Digunakan untuk konversi string ke bool.
- `strconv.FormatBool()` : Digunakan untuk konversi bool ke string.

## Konversi Data Menggunakan Teknik Casting

Keyword tipe data bisa digunakan untuk casting, atau konversi antar tipe data. Cara penggunaannya adalah dengan menuliskan tipe data tujuan casting sebagai fungsi, lalu menyisipkan data yang akan dikonversi sebagai parameter fungsi tersebut.

- Casting `string` ↔ `byte`
- Type Assertions Pada Interface Kosong (`interface{}`)
