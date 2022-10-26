# Time, Parsing Time, & Format Time

Pada bab ini kita akan belajar tentang pemanfaatan time, method-method yang disediakan, dan juga format & parsing data `string` ke `time.Time` dan sebaliknya.

Golang menyediakan package time yang berisikan banyak sekali komponen yang bisa digunakan untuk keperluan pemanfaatan waktu.

`Time disini maksudnya adalah gabungan date dan time, bukan hanya waktu saja.`

## Penggunaan `time.Time`

`time.Time` adalah tipe data untuk objek waktu. Ada 2 cara yang bisa digunakan untuk membuat data bertipe ini :

- Dengan menjadikan waktu sekarang sebagai time
- Membuat time dengan data ditentukan sendiri

Fungsi `time.Now()` mengembalikan objek `time.Time` dengan data adalah waktu sekarang. Bisa dilihat ketika di tampilkan, informasi yang muncul adalah sesuai dengan tanggal program tersebut dieksekusi.

Sedangkan `time.Date()` digunakan untuk membuat objek `time.Time` baru dengan informasi waktu yang bisa kita tentukan sendiri. Fungsi ini memiliki 8 buah parameter mandatory dengan skema bisa dilihat di kode berikut:
`time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)`

Objek cetakan fungsi `time.Now()`, informasi timezone-nya adalah relatif terhadap lokasi kita. Karena kebetulan saya berlokasi di Jawa Barat, Indonesia, maka akan terdeteksi masuk dalam GMT+7 atau WIB. Berbeda dengan objek cetakan fungsi `time.Date()` yang lokasinya bisa kita tentukan secara eksplisit misalnya UTC.

Selain menggunakan `time.UTC` untuk penentuan lokasi, tersedia juga `time.Local` yang nilainya adalah relatif terhadap waktu kita.

## Method Milik `time.Time`

Ada banyak method yang berhubungan dengan date, time, dan location yang tersedia pada tipe `time.Time`.

List lengkapnya bisa dilihat pada ebook Dasar Pemrograman GoLang by Noval Agung

## Parsing dari `string` ke `time.Time`

Data `string` bisa dikonversi menjadi `time.Time` dengan memanfaatkan `time.Parse`. Fungsi ini membutuhkan 2 parameter:

- Parameter ke-1 adalah layout format dari data waktu yang akan diparsing
- Parameter ke-2 adalah data string yang ingin diparsing

Layout format time di golang berbeda dibanding bahasa lain. Umumnya layout format yang digunakan adalah seperti `"DD/MM/YYYY"` dll, di Golang tidak.

Golang memiliki standar layout format yang cukup unik, contohnya seperti `2006-01-02 15:04:05` . Golang menggunakan `2006` untuk parsing tahun, bukan `YYYY`; `01` untuk parsing bulan; `02` untuk parsing hari; dan seterusnya.

Detailnya bisa dilihat pada ebook Dasar Pemrograman GoLang by Noval Agung

## Predefined Layout Format Untuk Keperluan Parsing Time

Golang juga menyediakan beberapa predefined layout format umum yang bisa dimanfaatkan. Jadi tidak perlu menuliskan kombinasi komponen-komponen layout format.

Salah satu predefined layout yang bisa digunakan adalah `time.RFC822`. Format ini sama dengan `02 Jan 06 15:04 MST`. Berikut contoh penerapannya :

```Go
var date, _ = time.Parse(time.RFC822, "14 Mar 99 05:00 WIB")
fmt.Println(date.String())
// 2015-09-02 08:00:00 +0700 WIB
```

Ada beberapa layout format lain yang bisa dimanfaatkan. Detailnya bisa dilihat pada ebook Dasar Pemrograman GoLang by Noval Agung

## Format `time.Time`

Setelah sebelumnya kita belajar tentang konversi `string` ke `time.Time`. Kali ini kita akan belajar sebaliknya, yaitu konversi `time.Time` ke `string`.

Method `Format()` digunakan untuk membentuk `string` menggunakan layout format tertentu. Layout format yang bisa digunakan adalah sama seperti pada parsing time.

## Handle Error Parsing `time.Time`

Ketika parsing `string` ke `time.Time`, sangat memungkinkan bisa terjadi error karena struktur data yang akan diparsing tidak sesuai layout format yang digunakan.

Untuk mendeteksi hal ini bisa memanfaatkan nilai kembalian ke-2 fungsi `time.Parse`. Berikut adalah contoh penerapannya.

```Go
var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
if err != nil {
  fmt.Println("error", err.Error())
  return
}
fmt.Println(date)
```

Atau

```Go
var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
if err != nil {
  fmt.Println("error", err.Error())
} else {
  fmt.Println(date)
}
```

Kode di atas menghasilkan error karena format tidak sesuai dengan skema data yang akan diparsing. Layout format yang harusnya digunakan adalah `06 Jan 15 03:04 MST`.
