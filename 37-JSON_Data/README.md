# JSON Data

**JSON** atau *Javascript Object Notation* adalah notasi standar yang umum digunakan untuk komunikasi data dalam web. JSON merupakan subset dari javascript.

Go menyediakan package `encoding/json` yang berisikan banyak fungsi untuk kebutuhan operasi json.

Pada chapter ini, kita akan belajar cara untuk konverstri string yang berbentuk json menjadi objek Go, dan sebaliknya.

## Decode JSON ke Variabel Object Struct

Di Go, data json dituliskan sebagai `string`. Dengan menggunakan `json.Unmarshal`, json string bisa dikonversi menjadi bentuk objek, entah itu dalam bentuk `map[string]interface{}` ataupun objek struct.

## Decode JSON ke `map[string]interface{}` & `interface{}`

Tak hanya ke objek cetakan struct, target decoding data json juga bisa berupa variabel bertipe `map[string]interface{}`.

Variabel bertipe `interface{}` juga bisa digunakan untuk menampung hasil decode. Dengan catatan pada pengaksesan nilai property, harus dilakukan casting terlebih dahulu ke `map[string]interface{}`.

### Note

Interface kosong atau empty interface yang dinotasikan dengan `interface{}` atau `any` (per go v1.18), merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun. Tipe data dengan konsep ini biasa disebut dengan dynamic typing.

`interface{}` merupakan tipe data, sehingga cara penggunaannya sama seperti pada tipe data lainnya, hanya saja nilai yang diisikan bisa apa saja. Contoh:

```Go
package main

import "fmt"

func main() {
    var secret interface{}

    secret = "ethan hunt"
    fmt.Println(secret)

    secret = []string{"apple", "manggo", "banana"}
    fmt.Println(secret)

    secret = 12.4
    fmt.Println(secret)
}
```

Keyword `interface` seperti yang kita tau, digunakan untuk pembuatan interface. Tetapi ketika ditambahkan kurung kurawal ({}) di belakang-nya (menjadi `interface{}`), maka kegunaannya akan berubah, yaitu sebagai tipe data.

Variabel bertipe `interface{}` bisa ditampilkan ke layar sebagai `string` dengan memanfaatkan fungsi print, seperti `fmt.Println()`. Tapi perlu diketahui bahwa nilai yang dimunculkan tersebut bukanlah nilai asli, melainkan bentuk string dari nilai aslinya.

Hal ini penting diketahui, karena untuk melakukan operasi yang membutuhkan nilai asli pada variabel yang bertipe `interface{}`, diperlukan casting ke tipe aslinya.

Misalnya, kita memiliki variabel `secret` bertipe `interface{}` yang digunakan untuk menampung nilai bertipe numerik. Ada kebutuhan untuk mengalikan nilai yang ditampung variabel tersebut dengan angka `10`. Maka perlu dilakukan casting ke tipe aslinya, yaitu `int`, setelahnya barulah nilai bisa dioperasikan. Proses casting dan perkaliannya dapat dilakukan secara bersamaan seperti kode berikut `number := secret.(int) * 10`. Jika kita hanya menuliskan `number := secret * 10` maka akan menimbulkan error dan program tidak bisa di compile.

## Decode Array JSON ke Array Object

Decode data dari array json ke slice/array objek masih sama, siapkan saja variabel penampung hasil decode dengan tipe slice struct.

## Encode Object ke JSON String

Setelah sebelumnya dijelaskan beberapa cara decode data dari json string ke objek, sekarang kita akan belajar cara **encode** data objek ke bentuk json string.

Fungsi `json.Marshal` digunakan untuk encoding data ke json string. Sumber data bisa berupa variabel objek cetakan `struct`, `map[string]interface{}`, atau `slice`.
