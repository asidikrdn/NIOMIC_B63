# Layout Format String

Di bab-bab sebelumnya kita telah banyak menggunakan layout format string seperti `%s`, `%d`, `%.2f`, dan lainnya; untuk keperluan menampilkan output ke layar ataupun untuk memformat string.

Layout format string digunakan pada konversi data ke bentuk string. Contohnya seperti `%.3f` yang untuk konversi nilai `double` ke `string` dengan 3 digit desimal.

Pada bab ini kita akan mempelajari satu per satu layout format string yang tersedia di Golang. Sampel data yang digunakan sebagai contoh adalah kode berikut.

```Go
type student struct {
  name string
  height float64
  age int32
  isGraduated bool
  hobbies []string
}
var data = student{
  name: "wick",
  height: 182.5,
  age: 26,
  isGraduated: false,
  hobbies: []string{"eating", "sleeping"},
}
```

## Layout Forrmat `%b`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 2 (biner).
Contoh :

```Go
fmt.Printf("%b\n", data.age) // 11010
```

## Layout Forrmat `%c`

Digunakan untuk memformat data numerik yang merupakan kode unicode, menjadi bentuk string karakter unicode-nya.
Contoh :

```Go
fmt.Printf("%c\n", 1400) // ո
fmt.Printf("%c\n", 1235) // a
```

## Layout Forrmat `%d`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 10 (basis bilangan yang kita gunakan).
Contoh :

```Go
fmt.Printf("%d\n", data.age) // 26
```

## Layout Forrmat `%e` atau `%E`

Digunakan untuk memformat data numerik desimal ke dalam bentuk notasi standar Scientific notation.
Contoh :

```Go
fmt.Printf("%e\n", data.height) // 1.825000e+02
fmt.Printf("%E\n", data.height) // 1.825000E+02
```

`1.825000E+02` maksudnya adalah `1.825 x 10^2`, dan hasil operasi tersebut adalah sesuai dengan data asli = `182.5`.

Perbedaan antara `%e` dan `%E` hanya huruf besar kecil karakter `e` pada hasil.

## Layout Forrmat `%f` atau `%F`

`%F` adalah alias dari `%f`. Keduanya memiliki fungsi yang sama. Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Secara default lebar digit desimal adalah 6 digit.
Contoh :

```Go
fmt.Printf("%f\n", data.height) // 182.500000
fmt.Printf("%.9f\n", data.height) // 182.500000000
fmt.Printf("%.2f\n", data.height) // 182.50
fmt.Printf("%.f\n", data.height) // 182
```

## Layout Forrmat `%g` atau `%G`

`%G` adalah alias dari `%g`. Keduanya memiliki fungsi yang sama. Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Lebar kapasitasnya sangat besar, pas digunakan untuk data yang jumlah digit desimalnya cukup banyak.

Bisa dilihat pada kode berikut perbandingan antara %e , %f , dan %g :

```Go
fmt.Printf("%e\n", 0.123123123123) // 1.231231e-01
fmt.Printf("%f\n", 0.123123123123) // 0.123123
fmt.Printf("%g\n", 0.123123123123) // 0.123123123123
```

Perbedaan lainnya adalah pada %g , lebar digit desimal adalah sesuai dengan datanya, tidak bisa dicustom seperti pada %f :

```Go
fmt.Printf("%g\n", 0.12) // 0.12
fmt.Printf("%.5g\n", 0.12) // 0.12
```

## Layout Forrmat `%o`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 8 (oktal).
Contoh :

```Go
fmt.Printf("%o\n", data.age) // 32
```

## Layout Forrmat `%p`

Digunakan untuk memformat data pointer, mengembalikan alamat pointer referensi variabelnya.

Alamat pointer dituliskan dalam bentuk numerik berbasis 16 dengan prefix `0x`.

Contoh :

```Go
fmt.Printf("%p\n", &data.name) // 0x2081be0c0
```

## Layout Forrmat `%q`

Digunakan untuk `escape` string. Meskipun string yang dipakai menggunakan literal `\` akan tetap di-escape.
Contoh :

```Go
fmt.Printf("%q\n",`" name \ height "`) // "\" name \\\\ height \""
```

## Layout Forrmat `%s`

Digunakan untuk memformat data string.
Contoh :

```Go
fmt.Printf("%s\n", data.name) // wick
```

## Layout Forrmat `%t`

Digunakan untuk memformat data boolean, menampilkan nilai bool -nya.
Contoh :

```Go
fmt.Printf("%t\n", data.isGraduated) // false
```

## Layout Forrmat `%T`

Berfungsi untuk mengambil tipe variabel yang akan diformat.
Contoh :

```Go
fmt.Printf("%T\n", data.name) // string
fmt.Printf("%T\n", data.height) // float64
fmt.Printf("%T\n", data.age) // int32
fmt.Printf("%T\n", data.isGraduated) // bool
fmt.Printf("%T\n", data.hobbies) // []string
```

## Layout Forrmat `%v`

Digunakan untuk memformat data apa saja (termasuk data bertipe `interface{}` ). Hasil kembaliannya adalah string nilai data aslinya.

Jika data adalah objek cetakan `struct`, maka akan ditampilkan semua secara property berurutan.
Contoh :

```Go
fmt.Printf("%v\n", data) // {wick 182.5 26 false [eating sleeping]}
```

## Layout Forrmat `%+v`

Digunakan untuk memformat struct, mengembalikan nama tiap property dan nilainya berurutan sesuai dengan struktur struct.
Contoh :

```Go
fmt.Printf("%+v\n", data)
// {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}
```

## Layout Forrmat `%#v`

Digunakan untuk memformat struct, mengembalikan nama dan nilai tiap property sesuai dengan struktur struct dan juga bagaimana objek tersebut dideklarasikan.
Contoh :

```Go
fmt.Printf("%#v\n", data)
// main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:[]string {"eating", "sleeping"}}
```

Ketika menampilkan objek yang deklarasinya adalah menggunakan teknik anonymous struct, maka akan muncul juga struktur anonymous struct nya.
Contoh :

```Go
var data = struct {
  name string
  height float64
}{
  name: "wick",
  height: 182.5,
}
fmt.Printf("%#v\n", data)
// struct { name string; height float64 }{name:"wick", height:182.5}
```

Format ini juga bisa digunakan untuk menampilkan tipe data lain, dan akan dimunculkan strukturnya juga.

## Layout Forrmat `%x` atau `%X`

Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 16 (heksadesimal).
Contoh :

```Go
fmt.Printf("%x\n", data.age)
// 1a
```

Jika digunakan pada tipe data string, maka akan mengembalikan kode heksadesimal tiap karakter.

```Go
var d = data.name

fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
// 7769636b

fmt.Printf("%x\n", d)
// 7769636b
```

`%x` dan `%X` memiliki fungsi yang sama. Perbedaannya adalah `%X` akan mengembalikan string dalam bentuk uppercase atau huruf kapital.

## Layout Format `%%`

Cara untuk menulis karakter `%` pada string format.

```Go
fmt.Printf("%%\n")
// %
```
