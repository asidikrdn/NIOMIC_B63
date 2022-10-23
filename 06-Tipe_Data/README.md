# Tipe Data pada GoLang

Golang mengenal beberapa jenis tipe data, diantaranya adalah tipe data numerik (desimal & non-desimal), string, dan boolean.

Di bab-bab sebelumnya secara tak sadar kita sudah menerapkan beberapa tipe data, seperti string dan tipe numerik int .

Pada bab ini, akan dijelaskan beberapa macam tipe data standar yang disediakan oleh Golang, dan bagaiman cara penggunaannya.

## Tipe Data Numerik Non-Desimal

Tipe data numerik non-desimal atau non floating point di Golang ada beberapa macam. Secara umum ada 2 tipe data yang perlu diketahui, yaitu:

- `uint` , merupakan tipe data untuk bilangan cacah (bilangan positif), dan
- `int` , merupakan tipe data untuk bilangan bulat (bilangan negatif dan positif)

Kedua tipe data di atas kemudian dibagi lagi menjadi beberapa, dengan pembagian berdasarkan lebar cakupan nilainya.

Detail cakupan nilai versi tipe data `int` dan `uint` :

- int8 (-128 s/d 127)
- int16 (-32768 s/d 32767)
- int32 (-2147483648 s/d 2147483647)
- int64 (-9223372036854775808 s/d 9223372036854775807)
- uint8 (0 s/d 255)
- uint16 (0 s/d 65535)
- uint32 (0 s/d 4294967295)
- uint64 (0 s/d 18446744073709551615)

Alias :

- byte (uint8)
- rune (int32)
- int (Minimal int32 {tergantung dari bit sistem operasi})
- uint (Minimal int32 {tergantung dari bit sistem operasi})

Dianjurkan untuk tidak sembarangan dalam menentukan tipe data variabel, sebisa mungkin tipe yang dipilih harus disesuaikan dengan nilainya, karena efeknya adalah ke alokasi memori variabel. Pemilihan tipe data yang tepat akan membuat pemakaian memori lebih optimal, tidak berlebihan.

Note : Template `%d` pada `fmt.Printf()` digunakan untuk memformat data numerik non-desimal.

## Tipe Data Numerik Desimal

Tipe data numerik desimal yang perlu diketahui ada 2, float32 dan float64 . Perbedaan kedua tipe data tersebut berada di lebar cakupan nilai desimal yang bisa ditampung. Untuk lebih jelasnya bisa merujuk ke spesifikasi IEEE-754 32-bit floating-point numbers.

Lebih detailnya versi tipe data float :

- float32 (1.18x10^-38 s/d 3.4x10^38)
- float64 (2.23x10^-308 s/d 1.80x10^308)
- complex64
- complex128

Contoh penggunaan tipe data ini bisa dilihat di kode berikut.

```Go
var decimalNumber = 2.62
fmt.Printf("bilangan desimal: %f\n", decimalNumber)
fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
```

Pada kode di atas, variabel decimalNumber akan memiliki tipe data float32 , karena nilainya berada di cakupan tipe data tersebut.

Template `%f` digunakan untuk memformat data numerik desimal menjadi string. Digit desimal yang akan dihasilkan adalah 6 digit. Pada contoh di atas, hasil format variabel decimalNumber adalah `2.620000`. Jumlah digit yang muncul bisa dikontrol menggunakan `%.nf` , tinggal ganti `n` dengan angka yang diinginkan. Contoh: `%.3f` maka akan menghasilkan 3 digit desimal, `%.10f` maka akan menghasilkan 10 digit desimal.

## Tipe Data bool (Boolean)

Tipe data `bool` berisikan hanya 2 variansi nilai, `true` dan `false` . Tipe data ini biasa dimanfaatkan dalam seleksi kondisi dan perulangan. Contoh sederhana penggunaan bool :

```Go
var exist bool = true
fmt.Printf("exist? %t \n", exist)
```

Gunakan `%t` untuk memformat data `bool` menggunakan fungsi `fmt.Printf()`.

## Tipe Data string

Ciri khas dari tipe data `string` adalah nilainya di apit oleh tanda quote atau petik dua (`"`). Contoh penerapannya:

```Go
var message string = "Halo"
fmt.Printf("message: %s \n", message)
```

Selain menggunakan tanda quote, deklarasi string juga bisa dengan tanda `grave accent/backticks`, tanda ini terletak di sebelah kiri tombol 1. Keistimewaan string yang dideklarasikan menggunakan backtics adalah membuat semua karakter didalamnya `tidak akan di escape`, termasuk `\n`, tanda petik dua dan tanda petik satu, baris baru, dan lainnya. Semua akan terdeteksi sebagai string. Berikut adalah contoh penerapannya.

```Go
var message = `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".`
fmt.Println(message)
```

Ketika dijalankan, output akan muncul sama persisi sesuai nilai variabel message di atas. Tanda petik dua akan muncul, baris baru juga muncul, sama persis.

## Nilai nil Dan Nilai Default Tipe Data

`nil` bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya `nil`, berarti variabel tersebut memiliki nilai kosong.

Semua tipe data yang sudah dibahas di atas memiliki nilai default. Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, akan ada nilai default-nya.

- Nilai default `string` adalah `""` (string kosong)
- Nilai default `bool` adalah `false`
- Nilai default tipe numerik non-desimal adalah `0`
- Nilai default tipe numerik desimal adalah `0.0`

`nil` adalah nilai kosong, benar-benar kosong. `nil` tidak bisa digunakan pada tipe data yang sudah dibahas di atas, karena kesemuanya sudah memiliki nilai default pada saat deklarasi. Ada beberapa tipe data yang bisa di-set nilainya dengan `nil`, diantaranya:

- pointer
- tipe data fungsi
- slice
- map
- channel
- interface kosong atau interface{}

Nantinya kita akan sering bertemu dengan nil ketika sudah masuk pada pembahasan bab-bab tersebut.
