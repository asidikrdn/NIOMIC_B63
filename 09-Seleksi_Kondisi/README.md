# Seleksi Kondisi pada GoLang

Seleksi kondisi digunakan untuk mengontrol alur program. Kalau dianalogikan, fungsinya mirip seperti rambu lalu lintas di jalan raya. Kapan kendaraan diperbolehkan melaju dan kapan harus berhenti, diatur oleh rambu tersebut. Sama seperti pada seleksi kondisi, kapan sebuah blok kode akan dieksekusi juga akan diatur.

Yang dijadikan acuan oleh seleksi kondisi adalah nilai bertipe bool , bisa berasa dari variabel, ataupun hasil operasi perbandingan. Nilai tersebut akan menentukan blok kode mana yang akan dieksekusi.

Golang memiliki 2 macam keyword untuk seleksi kondisi, yaitu `if else` dan `switch case`.

Note :
Golang tidak mendukung seleksi kondisi menggunakan ternary.
Statement seperti: `var data = (isExist ? "ada" : "tidak ada")` akan menghasilkan error.

Cara penulisannya bisa dilihat pada file `seleksi-kondisi.go`
