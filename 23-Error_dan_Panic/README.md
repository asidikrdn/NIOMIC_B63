# Error & Panic

Error merupakan topik yang penting dalam pemrograman golang. Di bagian ini kita akan belajar mengenai pemanfaatan error dan cara membuat custom error sendiri.

Kita juga akan belajar tentang penggunaan panic untuk menampilkan pesan error.

## Error

`error` merupakan sebuah tipe. Error memiliki beberapa property yang menampung informasi yang berhubungan dengan error yang bersangkutan.

Di golang, banyak sekali fungsi yang mengembalikan nilai balik lebih dari satu. Biasanya, salah satu kembalian adalah bertipe `error`. Contohnya seperti pada fungsi `strconv.Atoi()`.

`strconv.Atoi()` berguna untuk mengkonversi data string menjadi numerik. Fungsi ini mengembalikan 2 nilai balik. Nilai balik pertama adalah hasil konversi, dan nilai balik kedua adalah `error`.

Ketika konversi berjalan mulus, nilai balik kedua akan bernilai `nil`. Sedangkan ketika konversi gagal, kita bisa langsung tau penyebab error muncul dengan memanfaatkan nilai balik kedua.

Berikut merupakan contoh program sederhana untuk deteksi inputan dari user, apakah numerik atau bukan.

## Membuat Custom Error

Selain memanfaatkan error hasil kembalian fungsi, kita juga bisa membuat error sendiri dengan menggunakan fungsi `errors.New` (untuk menggunakannya harus import package `errors` terlebih dahulu).

Berikut merupakan contoh pembuatan custom error. Pertama siapkan fungsi dengan nama `validate()`, yang nantinya digunakan untuk pengecekan input, apakah inputan kosong atau tidak. Ketika kosong, maka error baru akan dibuat.

Selanjutnya di fungsi main, buat program sederhana untuk capture inputan user. Manfaatkan fungsi `validate()` untuk mengecek inputannya.

Fungsi `validate()` mengembalikan 2 data. Data pertama adalah nilai `bool` yang menandakan inputan apakah valid atau tidak. Data ke-2 adalah pesan error-nya (jika nputan tidak valid).

Fungsi `strings.TrimSpace()` digunakan untuk menghilangkan karakter spasi sebelum dan sesudah string. Ini dibutuhkan karena user bisa saja menginputkan spasi lalu enter.

Ketika inputan tidak valid, maka error baru dibuat dengan memanfaatkan fungsi `errors.New()`.

## Penggunaan Keyword `panic`

Panic digunakan untuk menampilkan trace error. Hasil keluarannya sama seperti `fmt.Println()` hanya saja informasi yang ditampilkan lebih detail.

Pada program yang telah kita buat tadi, ubah `fmt.Println()` yang berada di dalam blok kondisi else pada fungsi main menjadi `panic()`.

Ketika user menginputkan string kosong, maka error akan dimunculkan menggunakan fungsi `panic`.
