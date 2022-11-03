# Web Server

Go menyediakan package `net/http`, berisi berbagai macam fitur untuk keperluan pembuatan aplikasi berbasis web. Termasuk di dalamnya web server, routing, templating, dan lainnya.

Go memiliki web server sendiri, dan web server tersebut berada di dalam Go, tidak seperti bahasa lain yang servernya terpisah dan perlu diinstal sendiri (seperti PHP yang memerlukan Apache, .NET yang memerlukan IIS).

Pada chapter ini kita akan belajar cara pembuatan aplikasi web sederhana dan pemanfaatan template untuk mendesain view.

## Membuat Aplikasi Web Sederhana

Package `net/http` memiliki banyak sekali fungsi yang bisa dimanfaatkan. Di bagian ini kita akan mempelajari beberapa fungsi penting seperti routing dan start server.

Fungsi `http.HandleFunc()` digunakan untuk routing aplikasi web. Maksud dari routing adalah penentuan aksi ketika url tertentu diakses oleh user.

Pada kode contoh, 2 rute didaftarkan, yaitu `/` dan `/index`. Aksi dari rute `/` adalah menampilkan text `"halo"` di halaman website. Sedangkan `/index` menampilkan text `"apa kabar!"`.

Fungsi `http.HandleFunc()` memiliki 2 buah parameter yang harus diisi. Parameter pertama adalah rute yang diinginkan. Parameter kedua adalah callback atau aksi ketika rute tersebut diakses. Callback tersebut bertipe fungsi `func(w http.ResponseWriter, r *http.Request)`.

Pada pendaftaran rute `/index`, callback-nya adalah fungsi `index()`, hal seperti ini diperbolehkan asalkan tipe dari fungsi tersebut sesuai.

Fungsi `http.listenAndServe()` digunakan untuk menghidupkan server sekaligus menjalankan aplikasi menggunakan server tersebut. Di Go, 1 web aplikasi adalah 1 buah server berbeda.

Pada contoh, server dijalankan pada port `8080`.

Jalankan program dengan perintah `go run`. Perlu diingat, setiap ada perubahan pada file `.go`, `go run` harus dipanggil lagi.

Untuk menghentikan web server, tekan **CTRL+C** pada terminal atau CMD, di mana pengeksekusian aplikasi berlangsung.

## Penggunaan Template Web

Template engine memberikan kemudahan dalam mendesain tampilan view aplikasi website. Dan kabar baiknya Go menyediakan engine template sendiri, dengan banyak fitur yang tersedia di dalamnya.

Di sini kita akan belajar contoh sederhana penggunaan template untuk menampilkan data. Pertama siapkan dahulu template-nya. Selanjutnya ubah isi file `.go` seperti di contoh.

Fungsi `template.ParseFiles()` digunakan untuk parsing template, mengembalikan 2 data yaitu instance template-nya dan error (jika ada). Pemanggilan method `Execute()` akan membuat hasil parsing template ditampilkan ke layar web browser.

Pada kode contoh, variabel data disisipkan sebagai parameter ke-2 method `Execute()`. Isi dari variabel tersebut bisa diakses di-view dengan menggunakan notasi `{{.NAMA_PROPERTY}}` (nama variabel sendiri tidak perlu dituliskan, langsung nama property di dalamnya).

Pada contoh, statement di view `{{.Name}}` akan menampilkan isi dari `data.Name`.

### Note

Lokasi path tidak harus lengkap (ditelusuri dari folder root seperti `/etc` atau `Local Disk C`), jika project disimpan di `GOPATH` maka bisa langsung memanggil file sesuai hirarki path (seperti pada bahasa pemrograman lain).

Jika project tidak disimpan di `GOPATH`, maka kita harus menggunakan `GO MODULES`, cara ini adalah cara terbaru yang memungkinkan kita bisa menyimpan project secara bebas dimana saja. Untuk inisiasi modules, caranya masuk ke folder project > buka terminal > jalankan perintah `go mod init <nama_project>`.

Dengan menggunakan modules, ketika kita ingin mengakses sebuah file, cukup telusuri dari folder root project. Pada contoh ini, root projectnya adalah `NIOMIC-Golang`, maka ketika kita ingin mengakses file file di dalamnya, tinggal ditelusuri dari `NIOMIC-Golang`. Misalnya kita ingin mengakses file `template.html`, maka kita bisa mengaksesnya dengan `35-Web_Server/2-Web_Template/template.html`.
