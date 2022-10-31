# Regex

Regex/Regexp atau Regular Expression adalah suatu teknik yang digunakan untuk pencocokan string dengan pola tertentu. Regex bisa dimanfaatkan untuk pencarian dan pengubahan data string.

Go mengadopsi standar regex RE2, untuk melihat sintaks yang di-support engine ini bisa langsung merujuk ke dokumentasinya di [https://github.com/google/re2/wiki/Syntax].

## Penerapan `Regexp`

Fungsi `regexp.Compile()` digunakan untuk mengkompilasi ekspresi regex. Fungsi tersebut mengembalikan objek bertipe `regexp.*Regexp`.
Berikut contoh penerapan regex untuk pencarian karakter :

```Go
package main

import "fmt"
import "regexp"

func main() {
  var text = "banana burger soup"
  var regex, err = regexp.Compile(`[a-z]+`)

  if err != nil {
      fmt.Println(err.Error())
  }

  var res1 = regex.FindAllString(text, 2)
  fmt.Printf("%#v \n", res1)
  // []string{"banana", "burger"}

  var res2 = regex.FindAllString(text, -1)
  fmt.Printf("%#v \n", res2)
  // []string{"banana", "burger", "soup"}
}
```

Ekspresi `[a-z]+` maknanya adalah, semua string yang merupakan alphabet yang hurufnya kecil. Ekspresi tersebut di-compile oleh `regexp.Compile()` lalu disimpan ke variabel objek `regex` bertipe `regexp.*Regexp`.

Struct `regexp.Regexp` memiliki banyak method, salah satunya adalah `FindAllString()`, berfungsi untuk mencari semua string yang sesuai dengan ekspresi regex, dengan kembalian berupa slice string.

Jumlah hasil pencarian dari `regex.FindAllString()` bisa ditentukan. Contohnya pada `res1`, ditentukan maksimal `2` data saja pada nilai kembalian. Jika batas di set `-1`, maka akan mengembalikan semua data.

Ada cukup banyak method struct `regexp.*Regexp` yang bisa kita manfaatkan untuk keperluan pengelolaan string.

### Method `MatchString()`

Digunakan untuk mendeteksi apakah sebuah string memenuhi pola regexp tertentu. Methid ini akan mengembalikan nilai boolean `true`/`false`.
Format penulisan :
`varRegexHasilCompile.MatchString(textYangInginDiPeriksa)`

### Method `FindString()`

Digunakan untuk mencari sebuah string yang memenuhi pola regexp tertentu.
Fungsi ini hanya mengembalikan 1 buah hasil saja. Jika ada banyak substring yang sesuai dengan ekspresi regexp, akan dikembalikan yang pertama saja.
Format penulisan :
`varRegexHasilCompile.FindString(textUtama)`

### Method `FindStringIndex()`

Digunakan untuk mencari index string kembalian hasil dari operasi regexp.
Format penulisan :
`varRegexHasilCompile.FindStringIndex(textUtama)`

Outputnya berupa sebuah slice berisikan 2 angka, yaitu index awal string dan index akhir string.

### Method `FindAllString()`

Digunakan untuk mencari banyak string yang memenuhi kriteria regexp yang ditentukan.
Format penulisan :
`varRegexHasilCompile.FindAllString(textUtama, jumlahStringYangDiambil)`

### Method `ReplaceAllString()`

Digunakan untuk me-replace semua string yang memenuhi kriteria regexp tertentu dengan string lain.
Format penulisan :
`varRegexHasilCompile.ReplaceAllString(textUtama, stringPengganti)`

### Method `ReplaceAllStringFunc()`

Digunakan untuk me-replace semua string yang memenuhi kriteria regexp, dengan kondisi yang bisa ditentukan untuk setiap substring yang akan di replace.
Format penulisan :

```Go
varRegexHasilCompile.ReplaceAllStringFunc(textUtama, func(argPenampungTiapString string)string{
  if argPenampungTiapString == "burger" {
    return "potato"
  }
  return argPenampungTiapString
})
```

Pada contoh penulisan di atas, jika salah satu substring yang match adalah "burger" maka akan diganti dengan "potato", string selainnya tidak di replace.

### Method `Split()`

Digunakan untuk memisah string dengan pemisah adalah substring yang memenuhi kriteria regexp tertentu.

Jumlah karakter yang akan di split bisa ditentukan dengan mengisi parameter kedua fungsi regex.Split(). Jika di-isi -1 maka semua karakter yang memenuhi kriteria regex akan menjadi separator dalam operasi pemisahan/split. Contoh lain, jika di-isi 2, maka hanya 2 karakter pertama yang memenuhi kriteria regex akan menjadi separator dalam split tersebut.
Format penulisan :
`varRegexHasilCompile.Split(textUtama, jumlahKarakterYangAkanDiSplit)`
atau
`varRegexHasilCompile.Split(textUtama, jumlahElementHasilSplit)`
